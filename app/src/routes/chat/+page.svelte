<script lang='ts'>
    import { Avatar } from '@skeletonlabs/skeleton'
    import { messageFeed, rtc, username, socket, selectedPeer } from '$lib/stores.ts'
    import { RTC } from '$lib/rtc.ts'
    import { Packet } from '$lib/packet.ts'
    import { goto } from '$app/navigation'
    import { Peer } from '$lib/peer.ts'
    import { get } from 'svelte/store'


    if (localStorage.getItem('username') == null) {
        goto("../register")
    }

    async function sendOffer() {
        const offer: string = await rtc.makeOffer()
        let peer = get(selectedPeer)
        console.log(peer)
        socket.send(JSON.stringify(new Packet('post', 'passOffer', offer, $username, peer.name)) || '');
        console.log('offer sent')
    }

    async function sendAnswer(offer: string) {
        const answer: string = await rtc.makeAnswer(offer)
        let peer = get(selectedPeer)
        socket.send(JSON.stringify(new Packet('post', 'passAnswer', answer, $username, peer.name)) || '')
        console.log('answer sent')
    }
    
// ----------------------------------------------------------------


    socket.onopen = () => {
        console.log('WebSocket connection opened.')
        socket.send(JSON.stringify(new Packet('post', 'open', '', $username, 'server')) || '')
    }

    socket.onmessage = async (event) => {
        const response = JSON.parse(event.data)
        const payload = JSON.parse(response.payload)
        console.log(response)
        if (response.from) {
            rtc.from = response.from
        }
        switch (payload.type) {
            case 'offer':
                console.log('received offer', payload)
                await sendAnswer(JSON.stringify(payload))
                
            case 'answer':
                console.log('recieved answer', payload)
                await rtc.acceptAnswer(JSON.stringify(payload))
        }
    }

    socket.onclose = () => {
        console.log('WebSocket connection closed.')
    }

</script>


<div class='bg-transparent p-4 flex flex-col gap-3 overflow-y-auto'>
    
    {#if $selectedPeer}
    <button on:click={sendOffer} class="btn variant-filled-secondary">Offer</button>
    {/if}

    {#each $messageFeed as message }

    {#if message.name == $username}
    <div class="grid grid-cols-[auto_1fr] gap-2">
        <Avatar src="https://i.pravatar.cc/?img={message.avatar}" width="w-12" />
        <div class="card p-4 rounded-tl-none space-y-2 variant-soft-secondary">
            <header class="flex justify-between items-center">
                <p class="font-bold">{$username}</p>
                <small class="opacity-50">{message.timestamp}</small>
            </header>
            <p>{message.message}</p>
        </div>
    </div>
    {:else}
    <div class="grid grid-cols-[1fr_auto] gap-2">
        <div class="card p-4 rounded-tr-none space-y-2">
            <header class="flex justify-between items-center">
                <p class="font-bold">{message.name}</p>
                <small class="opacity-50">{message.timestamp}</small>
            </header>
            <p>{message.message}</p>
        </div>
        <Avatar src="https://i.pravatar.cc/?img={message.avatar}" width="w-12" />
    </div>
    {/if}
    {/each}
</div>
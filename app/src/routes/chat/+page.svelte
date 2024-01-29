<script lang='ts'>
    import { Avatar } from '@skeletonlabs/skeleton';
    import { messageFeed, rtc, usernameStore } from '$lib/stores.ts';
    import { RTC } from '$lib/rtc.ts'

    async function sendOffer() {
        const offer: string = await rtc.makeOffer()
        socket.send(offer)
        console.log('offer sent')
    }

    async function sendAnswer(offer: string) {
        const answer: string = await rtc.makeAnswer(offer)
        socket.send(answer)
        console.log('answer sent')
    }
    
    const socket = new WebSocket('ws://localhost:8080/ws')

    socket.onopen = () => {
        console.log('WebSocket connection opened.');
    };

    socket.onmessage = async (event) => {
        const response = JSON.parse(event.data)
        switch (response.type) {
            case 'offer':
                console.log('received offer')
                sendAnswer(JSON.stringify(response))
                break
            case 'answer':
                console.log('recieved answer')
                await rtc.acceptAnswer(JSON.stringify(response))
                break
        }
    };

    socket.onclose = () => {
        console.log('WebSocket connection closed.');
    };

</script>


<div class='bg-transparent p-4 flex flex-col gap-3 overflow-y-auto'>
    <button on:click={sendOffer} class="btn variant-filled-secondary">Offer</button>
    {#each $messageFeed as message }

    {#if message.name == $usernameStore}
    <div class="grid grid-cols-[auto_1fr] gap-2">
        <Avatar src="https://i.pravatar.cc/?img={message.avatar}" width="w-12" />
        <div class="card p-4 variant-soft rounded-tl-none space-y-2">
            <header class="flex justify-between items-center">
                <p class="font-bold">{$usernameStore}</p>
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
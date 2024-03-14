import { Packet } from '$lib/packet.ts'
import { rtc, selectedPeer, username, savedPeers } from '$lib/stores.ts'
import { get } from 'svelte/store'
import { Peer } from '$lib/peer.ts'


const socket = new WebSocket('ws://176.230.36.233:27357/ws')

socket.onopen = () => {
    console.log('WebSocket connection opened.')
    let packet: Packet = new Packet('connect', localStorage.getItem('publicKey') || '', get(username) || 'guest', 'server')
    sendPacket(packet)
}

socket.onmessage = async (event) => {
    const response = JSON.parse(event.data)
    const payload = JSON.parse(response.payload)
    console.log(response)
    if (response.from) {
        rtc.from = response.from

        let newPeer = new Peer(response.from, response.from)
		selectedPeer.set(newPeer)
		
		let peers = get(savedPeers)
		const peerExists = peers.some((peer) => peer.name === newPeer.name)

		if (!peerExists) {
			// Peer doesn't exist, add it to savedPeers
			savedPeers.update((peers) => [...peers, newPeer]) // Adjust update function according to your store implementation
			localStorage.setItem('savedPeers', JSON.stringify(peers))
		}

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

async function sendAnswer(offer: string) {
    const answer: string = await rtc.makeAnswer(offer)
    let peer = get(selectedPeer)
    let packet = new Packet('passPacket', answer, get(username) || '', peer.name)

    sendPacket(packet)
    console.log('answer sent')
}

export function sendPacket(packet: Packet) {
    socket.send(JSON.stringify(packet))
}
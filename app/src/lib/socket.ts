import { Packet } from '$lib/packet.ts'
import { rtc, selectedPeer, username } from '$lib/stores.ts'
import { get } from 'svelte/store'

const socket = new WebSocket('ws://87.68.161.205:27357/ws')

socket.onopen = () => {
    console.log('WebSocket connection opened.')
    let packet: Packet = new Packet('post', 'open', '', get(username) || '', 'server')
    sendPacket(packet)
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

async function sendAnswer(offer: string) {
    const answer: string = await rtc.makeAnswer(offer)
    let peer = get(selectedPeer)
    let packet = new Packet('post', 'passAnswer', answer, get(username) || '', peer.name)

    sendPacket(packet)
    console.log('answer sent')
}

export function sendPacket(packet: Packet) {
    socket.send(JSON.stringify(packet))
}
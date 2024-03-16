import { Packet } from '$lib/packet.ts'
import { selectedPeer, username, savedPeers, rtcs } from '$lib/stores.ts'
import { get } from 'svelte/store'
import { Peer } from '$lib/peer.ts'
import { RTC } from './rtc'

const socket = new WebSocket('ws://176.230.36.233:27357/ws')
// TODO: fetch this IP from somewhere

socket.onopen = () => {
    console.log(`connected to ${socket.url} as @${get(username)}.`)
    sendPacket(new Packet('connect', localStorage.getItem('publicKey') || 'TODO: error', get(username) || 'guest', 'server'))
}

socket.onmessage = async (event) => {
    const packet: Packet = JSON.parse(event.data)
    const payload: RTCSessionDescription = JSON.parse(packet.payload)
    console.log("packet from server:", packet)

    const newPeer = new Peer(packet.from, packet.from)

    // check if peer needs to be added to the peer list
    const peerExists = get(savedPeers).some((peer) => peer.name === newPeer.name)
    if (!peerExists) {
        // add it to savedPeers
        savedPeers.update(peers => [...peers, newPeer])
    }

    selectedPeer.set(newPeer)

    switch (packet.action) {
        case 'passPacket':
            switch (payload.type) {
                case 'offer':
                    console.log('received offer:', payload)
                    let newRtc = new RTC(packet.from)
                    rtcs.update(rtcs => [...rtcs, newRtc])
                    newRtc.sendAnswer(JSON.stringify(payload))   
            
                    break
                    
                case 'answer':
                    console.log('recieved answer', payload)
                    get(rtcs).forEach(rtc => {
                        if (rtc.peerName == get(selectedPeer).name) {
                            rtc.acceptAnswer(JSON.stringify(payload))
                            return
                        }
                    });
                break
            }
            break
        default:
           alert('faulty packet action' + packet.action)
    }
}

socket.onclose = () => {
    alert('Disconnected from the server.')
}

export function sendPacket(packet: Packet) {
    socket.send(JSON.stringify(packet))
}
import { Packet } from '$lib/packet'
import { selectedPeer, username, savedPeers, rtcs } from '$lib/stores'
import { get } from 'svelte/store'
import { Peer } from '$lib/peer'
import { RTC } from './rtc'

async function fetchServer(): Promise<string> {
    const response = await fetch('/api/servers')
    const servers = JSON.parse(await response.text());
    const randomIndex = Math.floor(Math.random() * servers.length);
    return servers[randomIndex].ip;
}

const socket = new WebSocket('ws://' + await fetchServer() + ':27357/ws')

socket.onopen = () => {
    console.log(`listening to ${socket.url}.`)
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
                    console.log('received offer from:', packet.from)
                    let newRtc = new RTC(packet.from)
                    rtcs.update(rtcs => [...rtcs, newRtc])
                    newRtc.sendAnswer(JSON.stringify(payload))
                    break

                case 'answer':
                    console.log('recieved answer from:', packet.from)
                    get(rtcs).forEach(rtc => {
                        if (rtc.peerName == get(selectedPeer).name) {
                            rtc.acceptAnswer(JSON.stringify(payload))
                            return
                        }
                    })
                    break
            }
            break
        default:
            alert('bad packet action' + packet.action)
    }
}

socket.onclose = () => {
    alert('Disconnected from the server.')
}

export function sendPacket(packet: Packet) {
    socket.send(JSON.stringify(packet))
}
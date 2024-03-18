import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import { Peer } from '$lib/peer.ts'
import type { Bubble } from '$lib/types.ts'
import type { RTC } from './rtc'

function getPeers(): Peer[] {
    let json = JSON.parse(localStorage.getItem('savedPeers') || '[]')
    let peers: Peer[] = []
    json.forEach((element: { name: string; avatar: string }) => {
        peers.push(new Peer(element.name, element.avatar))
    })
    return peers
}

export const savedPeers: Writable<Peer[]> = writable(getPeers())
savedPeers.subscribe((value) => {
    localStorage.setItem('savedPeers', JSON.stringify(value));
})

export const selectedPeer: Writable<Peer> = writable()

export const messageFeed: Writable<Bubble[]> = writable([])

export const rtcs: Writable<RTC[]> = writable([])

export const username = writable(localStorage.getItem('username') || '')

export const connectionStatus: Writable<string> = writable('closed')

export const privateKey: string = localStorage.getItem('privateKey') || ''
export const publicKey: string = localStorage.getItem('publicKey') || ''

export const bio: string  = localStorage.getItem('bio') || ''
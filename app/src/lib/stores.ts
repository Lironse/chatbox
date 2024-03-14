import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import { RTC } from '$lib/rtc.ts'
import { Peer } from '$lib/peer.ts'
import type { Bubble } from '$lib/types.ts'

export const selectedPeer: Writable<Peer> = writable()
export const savedPeers: Writable<Peer[]> = writable(JSON.parse(localStorage.getItem('savedPeers') || '[]'))

export const messageFeed: Writable<Bubble[]> = writable([])

export const username = writable(localStorage.getItem('username') || '')

export const rtc = new RTC()
export const connectionStatus: Writable<string> = writable()

export const privateKey: string = localStorage.getItem('privateKey') || ''
export const publicKey: string = localStorage.getItem('publicKey') || ''
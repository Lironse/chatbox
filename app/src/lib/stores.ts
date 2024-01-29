import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store'
import { RTC } from '$lib/rtc.ts'

type Bubble = {
    avatar: string;
    name: string;
    timestamp: string;
    message: string;
};

export const selectedPeer = writable('');
export const savedPeers = writable(localStorage.getItem('savedPeers') ? JSON.parse(localStorage.getItem('savedPeers') || '')   : []);
export const messageFeed: Writable<Bubble[]> = writable([]);
export const usernameStore = writable(localStorage.getItem('username') ? localStorage.getItem('username')  : '');
export const rtc = new RTC()
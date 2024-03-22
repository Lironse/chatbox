import { messageFeed } from '$lib/stores.ts'
import type { Bubble } from './types';

export function addMessageToChat(content: string, name: string, avatar: string) {
    const now = new Date();
    const hours = now.getHours().toString().padStart(2, '0');
    const minutes = now.getMinutes().toString().padStart(2, '0');
    const time = `${hours}:${minutes}`;

    const msg: Bubble = {
        avatar: avatar,
        name: name,
        timestamp: time,
        message: content
    };

    messageFeed.update((messages) => [...messages, msg]);

}
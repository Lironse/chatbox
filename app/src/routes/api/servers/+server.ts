import { json } from "@sveltejs/kit";
import { list } from '$lib/serverlist.ts'
import type { ServerEntry } from "$lib/types.js";

export function GET() {
    return json(list)
}

export async function POST(requestEvent) {
    const { request } = requestEvent;

    const { id, ip } = await request.json();
    const server: ServerEntry = {
        id: id,
        ip: ip
    }
    if (list.includes(server)) {
        return json("This IP is already registered")
    }

    list.push(server);
    return json(list, { status: 201 });
}
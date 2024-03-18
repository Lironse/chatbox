import { json } from "@sveltejs/kit";
import { list } from '$lib/serverlist.ts'

export function GET() {
    return json(list)
}

export async function POST(requestEvent) {
    const { request } = requestEvent;
    const { text } = await request.json();
    list.push(text)
    return json(list, {status: 201});
}
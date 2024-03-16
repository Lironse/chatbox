<script lang="ts">
	import { rtcs, username, selectedPeer } from '$lib/stores.ts';
	import { addMessageToChat } from '$lib/index.ts';
	import { get } from 'svelte/store';

	let addMessageInput: string;

	function addMessage() {
		addMessageToChat(addMessageInput, get(username), get(username));
		get(rtcs).forEach((rtc) => {
			if (rtc.peerName == get(selectedPeer).name) {
				rtc.sendMessage(addMessageInput);
				return;
			}
		});
		addMessageInput = '';
	}
</script>

<form class="p-4" name="addPeer" on:submit={addMessage}>
	<input
		id="message-input"
		bind:value={addMessageInput}
		class="input"
		type="text"
		placeholder="Enter message"
		maxlength="50"
		autocomplete="off"
		spellcheck="false"
	/>
	<input type="submit" hidden />
</form>

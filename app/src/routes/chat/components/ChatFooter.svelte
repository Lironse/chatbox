<script lang="ts">
	import { messageFeed, usernameStore, rtc } from '$lib/stores.ts';
    import { addMessageToChat } from '$lib/index.ts';

	let addMessageInput: string;
    
    function addMessage() {
        addMessageToChat(addMessageInput, $usernameStore, '1')
        rtc.sendMessage(addMessageInput)
        addMessageInput = ''
    }

</script>

<form class="p-4" name="addPeer" on:submit={addMessage}>
	<input
		bind:value={addMessageInput}
		class="input"
		type="text"
		placeholder="Enter message"
		maxlength="50"
	/>
	<input type="submit" hidden />
</form>

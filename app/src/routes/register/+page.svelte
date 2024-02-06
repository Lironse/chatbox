<script lang='ts'>
	import InfoCard from './components/InfoCard.svelte'
	import RegisterForm from './components/RegisterForm.svelte'
	
	let registered = false

	function waitForRegistration(value) {
   		registered = value
  	}

</script>

<div class="grid h-full">
	{#if !registered}
		<RegisterForm {waitForRegistration}/>
	{:else}
		<InfoCard />
	{/if}
</div>

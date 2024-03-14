<script>
    import { Avatar } from '@skeletonlabs/skeleton'
    import { username } from '$lib/stores.ts'

</script>

<div class='flex flex-row gap-5 hover:card p-2'>
    <Avatar src='https://i.pravatar.cc/?img={$username}' />
    <div class='flex flex-col'>
        <h2 class='h2 font-bold'>{$username}</h2>
        bio example...
    </div>
</div>
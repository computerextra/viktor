<script>
	import { page } from '$app/state';
	export let href;
	export let title;
	export let path;
</script>

<a
	class="focus:outline-hidden font-medium text-gray-600 hover:text-gray-400 focus:text-gray-400 dark:text-neutral-400 dark:hover:text-blue-800 dark:focus:text-blue-800"
	{href}
	class:active={path == href}
>
	{title}
</a>

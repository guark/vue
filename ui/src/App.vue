<template>
	<div id="app">
		<HelloWorld msg="Hello World Guark Vue Template!" />
	</div>
</template>

<style lang="scss">
body, html {
	width: 100%;
	height: 100%;
	padding: 0;
	margin: 0;
	border: 0;
}

#app {
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	height: 100%;
	width: 100%;
	display: flex;
	justify-content: center;
	align-items: center;
	color: rgba(247, 250, 252, 0.8);
	background-image: linear-gradient(to top, #fad0c4 0%, #ffd1ff 100%);
}

#nav {
	padding: 30px;

	a {
		font-weight: bold;
		color: #2c3e50;

		&.router-link-exact-active {
			color: #42b983;
		}
	}
}

// Disable image drag and drop.
img {
	-webkit-user-drag: none;
	-khtml-user-drag: none;
	-moz-user-drag: none;
	-o-user-drag: none;
	user-drag: none;
}

* {
	user-select: none;
	-moz-user-select: none;
	-webkit-user-select: none;
	-ms-user-select: none;
	touch-callout: none;
	-webkit-touch-callout: none;
}

input {
	user-select: auto;
	-moz-user-select: auto;
	-webkit-user-select: auto;
	-ms-user-select: auto;
}
</style>

<script>
import g from "guark"
import HelloWorld from '@/components/HelloWorld.vue'

export default {
	name: "App",

	components: {
		HelloWorld,
	},

	created()
	{
		g.env().then(env =>
		{
			// Disable right click menu on production, you can implement your own!
			if (!!env.dev_mode) {
				document.addEventListener('contextmenu', e => e.preventDefault());
			}
		})
	},
}
</script>
<template>
	<div class="hello">
		<h1 style="color:#5556e4;">{{ msg }}</h1>
		<img style="width: 300px; border-radius: 5%;margin-bottom: 20px" :src="require('@/assets/icon.png')">
		<div>
			<button class="btn" @click="gcall()">Call Go Function!</button>
			<button class="btn" @click="clipboardWrite()">Call Plugin (Write "guark" to clipboard)!</button>
			<button class="btn" @click="error()">Call Dialog error!</button>
			<button class="btn" @click="info()">Call Plugin Dialog info!</button>
			<button class="btn" @click="notify()">Call Plugin Notify!</button>
			<button class="btn" @click="select()">Call Plugin Dialog Select file!</button>
		</div>
		<div style="padding: 10px;color: rgba(0, 0, 0, 0.5)">
			AppName: {{ env.app_name }}, DevMode: {{ env.dev_mode }}, Version: {{ env.app_version }}
		</div>
	</div>
</template>

<script>
import g from 'guark'

export default
{
	name: 'HelloWorld',

	props: {
		msg: String
	},

	data()
	{
		return {
			env: {},
		}
	},

	created()
	{
		g.env().then(env => this.env = env)
	},

	methods:
	{
		async gcall()
		{
			// You can await too.
			var hello_data = await g.call("hello_world", { name: "Guark working?" })

			console.log(`hello_world call return data: ${hello_data}`)
		},

		clipboardWrite()
		{
			g.call("clipboard.write", { text: "guark" }).then(() =>
			{
				g.call("notify.send", {message: "Guark copied to your clipboard"})
			})
		},

		error()
		{
			g.call("dialog.error", { message: "this is a error from js api! is this working?" })
		},

		info()
		{
			g.call("dialog.info", { message: "this is a info from js api! is this working?" })
		},

		notify()
		{
			g.call("notify.send", { message: "This is a notify test call from ui javascript api." })
		},

		select()
		{
			g.call("dialog.file", {title: "Select File For Guark APP!" }).then(res =>
			{
				g.call("notify.send", {message: "File Selected: "+ res})
			})
		}
	}
}
</script>

<style>
h3 {
	margin: 40px 0 0;
}

ul {
	list-style-type: none;
	padding: 0;
}

li {
	display: inline-block;
	margin: 0 10px;
}

a {
	color: #42b983;
}

.btn {
	margin: 3px;
	background: #5556e4;
	padding: 12px;
	border-radius: 10px;
	cursor: pointer;
	border: none;
	outline: none;
	font-size: 15px;
	font-weight: 500;
	color: #FFF;
	box-shadow: 0pt 1px 4px 1px rgba(85, 86, 228,0.5);
}
</style>

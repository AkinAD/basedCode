<template>
    <header class="header">
        <h1>Smart Shopping™ </h1>
        <Navbar/>
    </header>
</template>

<script>
import Navbar from './navbar.vue'
export default {
    name: "Header",
    components: {
        Navbar
    }
}
</script>

<style scoped>
    .header {
        background: #333;
        color: #fff;
        text-align: center;
        padding: 5px;
    }
</style>
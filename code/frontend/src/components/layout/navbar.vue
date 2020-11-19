<!--
Possible Navbar source:
https://vuejsexamples.com/a-simple-and-pretty-navbar-for-your-vue-projects/
-->
<template>
    <header class="navbar">
        <div id="nav"> <!--temp navbar-->
            <router-link to="/">Home</router-link>
            <router-link to="/account">Account</router-link> 
            <router-link to="/browse">Browse</router-link> 
            <router-link to="/signin">Signin</router-link> 
            <router-link to="/signup">Signup</router-link> 
        </div>
    </header>
</template>

<style scoped>
    .navbar a {
        color: #fff;
        padding-right: 10px;
        text-decoration: none;
    }
</style>
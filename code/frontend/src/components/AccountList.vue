<template>
    <div>
        <h3> Example Account List  </h3>
        <div class = "accounts">
            <div v-for="account in allAccounts" :key="account.id" class="account">
                {{account.name}}
            </div>
        </div>
    </div>
</template>

<script>
import {mapGetters, mapActions} from 'vuex'
export default {
    name: "Accounts",
    methods: {
        ...mapActions(['fetchAccounts'])
    },
    computed: mapGetters(['allAccounts']),
    created() {
        this.fetchAccounts();
    }
}
</script>

<style>
    .accounts {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-gap: 1rem;
    }
    .account {
        border: 1px solid #ccc;
        background: #4eb4f8;
        padding: 1rem;
        border-radius: 5px;
        text-align: center;
        color: black;
        position: relative;
        cursor: pointer;
    }
</style>
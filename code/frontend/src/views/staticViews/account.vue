<template>
    <div class="container">
        <AccountList/>
    </div>
</template>

<script>
import AccountList from '../../components/AccountList'
export default {
  name: 'Home',
  components: {
    AccountList
  }
}
</script>

<style>
body {
    font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
    line-height: 1.6;
    background: #e8f4f7;
  }
  .container {
    max-width: 1100px;
    margin: auto;
    overflow: auto;
    padding: 0 2rem;
  }
</style>

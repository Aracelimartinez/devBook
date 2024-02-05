<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const redirectToLogin = () => {
router.push('/login');
};

</script>

<template>
  <main>
    <div id="register-page" class="h-dvh text-color-text-body">

      <div class="w-10/12 flex flex-col">
        <img class="m-5 w-6/12" alt="LogIn image"  src="@/assets/images/undraw_fingerprint_login.svg"/>
        <h1 class="ml-5 mb-5 text-3xl font-bold text-left">Junte-se hoje à comunidade do <span class="text-color-primary">DevBook</span></h1>
      </div>
        <!-- Register card -->
      <div class=" flex flex-col p-7 text-base register-card w-10/12">
        <h2 class="mb-3 text-xl font-bold">Crie a sua conta!</h2>
        <form class="flex flex-col" action="" name="login-fom">
          <input class="my-3 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="text" placeholder="&#xf5b7;  Nome completo" required="required">
          <input class="my-3 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="text" placeholder="&#xf007;  Username" required="required">
          <input class="my-3 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="email" placeholder="&#xf0e0;  E-mail" required="required">
          <input class="my-3 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="password" placeholder="&#xf084;  Senha" required="required">
          <input class="my-3 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="password" placeholder="&#xf084;  Confirmar senha" required="required">
          <button type="submit" class="mt-3 py-2 px-6  text-white bg-color-primary rounded-md hover:bg-color-secondary duration-150">Registrar-se</button>
        </form>
        <div class="my-3 text-line">
          <span>ou</span>
        </div>
        <button @click="redirectToLogin" class="mb-3 py-2 px-6 bg-color-details rounded-md hover:opacity-80 duration-150">Faça seu login</button>
      </div>

    </div>
  </main>
</template>

<style lang="scss" scoped>
  #register-page {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 16px;
    place-items: center;
  }

  .register-card {
    background-color: #fff;
    box-shadow: 0px 0px 2px 2px rgba(0, 0, 0, 0.075);
    border-radius: 8px;
    max-width: 600px;
  }

  .text-line {
  position: relative;
  display: inline-block;
  text-align: center;
  }
  .text-line::before, .text-line::after {
    content: "";
    position: absolute;
    top: 0.65em;
    width: 40%;
    height: 0.05em;
    background-color: #666666;
    left: 0em;
  }
  .text-line::after {
    left: initial;
    right: 0em;
  }

  .fontAwesome {
  font-family: 'Roboto', FontAwesome, sans-serif;
  font-weight: 100;
  font-size: 14px;
  }

  // Smallest device
  @media (min-width: 100px) and (max-width: 575px) {
    #register-page {
      grid-template-columns: 1fr;
    }
  }

  // Small devices (landscape phones, 576px and up)
  @media (min-width: 576px) and (max-width: 767px){
    #register-page {
      grid-template-columns: 1fr;
    }
  }

  // Medium devices (tablets, 768px and up)
  @media (min-width: 768px) {
    #register-page {
      grid-template-columns: 1fr 1fr;
      margin: 0px 80px;
    }
  }
</style>

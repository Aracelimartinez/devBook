<script setup>
  import { useRouter } from 'vue-router';

  const router = useRouter();

  const redirectToRegister = () => {
  router.push('/register');
  };
</script>

<template>
  <main>
    <div id="login-page" class="h-dvh text-color-text-body">
      <img class="m-5 p-5 w-10/12" alt="LogIn image"  src="@/assets/images/undraw_login.svg"/>

      <div class="w-10/12 flex flex-col justify-center">
        <h1 class="mb-5 text-3xl font-bold text-center">Seja bem vindo ao <span class="text-color-primary">DevBook</span></h1>
        <!-- Sign in card -->
        <div class=" flex flex-col p-7 text-base signin-card">
          <h2 class="mb-3 text-xl font-bold">Iniciar sessão</h2>
          <form class="flex flex-col" action="" name="login-fom">
            <input class="my-3 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="email" placeholder="&#xf0e0;  E-mail" required="required">
            <input class="my-3 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="password" placeholder="&#xf084;  Senha" required="required">
            <button type="submit" class="mt-3 py-2 px-6  text-white bg-color-primary rounded-md hover:bg-color-secondary duration-150">Entrar</button>
          </form>
          <div class="my-3 text-line">
            <span>ou</span>
          </div>
          <button @click="redirectToRegister" class="mb-3 py-2 px-6 bg-color-details rounded-md hover:opacity-80 duration-150">Criar conta</button>
        </div>
      </div>
    </div>
  </main>
</template>

<style lang="scss" scoped>
  #login-page {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 16px;
    place-items: center;
  }

  .signin-card {
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
  }

  // Smallest device
  @media (min-width: 100px) and (max-width: 575px) {
    #login-page {
      grid-template-columns: 1fr;
    }
  }

  // Small devices (landscape phones, 576px and up)
  @media (min-width: 576px) and (max-width: 767px){
    #login-page {
      grid-template-columns: 1fr;
    }
  }

  // Medium devices (tablets, 768px and up)
  @media (min-width: 768px) {
    #login-page {
      grid-template-columns: 1fr 1fr;
      margin: 0px 80px;
    }
  }
</style>

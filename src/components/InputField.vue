<template>
    <label for=""></label>
    <input type="text" placeholder="" name="">
</template>

<script>
export default {
    name: 'InputField.vue',
    
    data() {
        return {
            inputs: [
                {
                    name: "email",
                    placeholder: "E-mail"
                },
                {
                    name: "password",
                    placeholder: "Password"
                },
                {
                    name: "brand",
                    placeholder: "Brand (ck, guess, levis, bershka)"
                },
                {
                    name: "type",
                    placeholder: "Type (Coat, Jacket, Denim, Leather)"
                }
            ]
        }
    }
}
</script>
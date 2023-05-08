<template>
    <TheHeader />
    <main>
        <h1>{{email}} page</h1>
        <div v-for="item in it_wardrobe_item" :key="item.id">
            <div>
                <p>{{ item.item_type }}</p>
                <p>{{ item.brand_name }}</p>
                <p>{{ item.size }}</p>
            </div>
            <div v-if="item.size_type == it_garment_size_type.name && item.size == it_garment_size_type_item.name">
                <p>{{ it_garment_size_type_item.name }}</p>
            </div>
        </div>
    </main>
</template>
<script>
import TheHeader from './TheHeader.vue'
import axios from 'axios'
export default {
    name: 'TheHome',

    data() {
        return {
            email:'',
            it_wardrobe_item:[],
            it_garment_size_type:[],
            it_garment_size_type_item:[]
        }
    },

    components: {
        TheHeader
    },

    async mounted() {
        let user = localStorage.getItem('user-info')
        this.email = JSON.parse(user).email
        if(!user) {
            this.$router.push({name:'SignUp'})
        }

        let result = await axios.get("http://localhost:3000/it_wardrobe_item")
        let sizeType = await axios.get("http://localhost:3000/it_garment_size_type")
        let sizes = await axios.get("http://localhost:3000/it_garment_size_type_item")
        this.it_wardrobe_item = result.data
        this.it_garment_size_type = sizeType.data
        this.it_garment_size_type_item = sizes.data
    }
}
</script>
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
            <!--div v-if="item.size_type == it_garment_size_type.name && item.size == it_garment_size_type_item.name">
                <p>{{ it_garment_size_type_item.name }}</p>
            </div-->
        </div>
        <div>
            <form class="add">
                <input type="text" name="item_type" placeholder="Type" v-model="wardrobe_item.item_type">
                <input type="text" name="brand_name" placeholder="Brand" v-model="wardrobe_item.brand_name">
                <input type="text" name="size" placeholder="Size" v-model="wardrobe_item.size">
                <button type="button" v-on:click="addWardrobeItem">Add</button>
            </form>
        </div>
    </main>
</template>
<script>
import TheHeader from './TheHeader.vue'
import axios from 'axios'
export default {
    name: 'TheHome',

    components: {
        TheHeader
    },

    data() {
        return {
            email:'',
            it_wardrobe_item:[],
            it_garment_size_type:[],
            it_garment_size_type_item:[],
            wardrobe_item: {
                item_type:'',
                brand_name:'',
                size:''
            }
        }
    },

    methods: {
        async addWardrobeItem() {
            const result = await axios.post("http://localhost:3000/it_wardrobe_item", {
                item_type:this.wardrobe_item.item_type,
                brand_name:this.wardrobe_item.brand_name,
                size:this.wardrobe_item.size
            })
            if(result.status == 201) {
                //this.$router.push({name:'TheLogin'})
                //this.$router.replace(this.$route.path)
                this.$router.go(0)
                //this.$forceUpdate()
            }
            console.log(result)
        }
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
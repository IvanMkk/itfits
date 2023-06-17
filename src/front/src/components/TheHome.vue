<template>
    <TheHeader />
    <main>
        <h1>{{email}} page</h1>
        <!--div v-for="item in it_wardrobe_item" :key="item.id" class="card">
            <div class="data item">
                <p class="garment-type">{{ item.item_type }}</p>
                <p class="garment-brand">{{ item.brand_name }}</p>
                <p class="garment-size">{{ item.size }}</p>
            </div-->
            <!--div v-if="item.size_type == it_garment_size_type.name && item.size == it_garment_size_type_item.name" class="data sizes">
                <p>{{ it_garment_size_type_item.name }}</p>
            </div-->
        <!--/div-->
        <div>
            <form class="form add">
                <div class="form__group">
                    <input type="text" id="brand" name="brand" class="form__control" placeholder="Brand (ck, guess, levis, bershka)" v-model="wardrobe_item.brand_name">
                    <label for="brand" class="form__label none">Brand (ck, guess, levis, bershka)</label>
                </div>
                <div class="form__group">
                    <input type="text" id="type" name="type" class="form__control" placeholder="Type (Coat, Jacket, Denim, Leather)" v-model="wardrobe_item.item_type">
                    <label for="type" class="form__label none">Type (Coat, Jacket, Denim, Leather)</label>
                </div>
                <div class="form__group">
                    <input type="text" id="size" name="size" class="form__control" placeholder="Size" v-model="wardrobe_item.size">
                    <label for="type" class="form__label none">Size</label>
                </div>
                <div class="buttons">
                    <button type="button" v-on:click="addWardrobeItem" class="button-small none" id="btn-add">Add</button>
                    <button class="button-second" id="btn-cancel">Cancel</button>
                </div>
            </form>
        </div>
    </main>
</template>
<script>
import TheHeader from './TheHeader.vue'
//import axios from 'axios'
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
        /*async addWardrobeItem() {
            const result = await axios.post("http://0.0.0.0:3000/v1/it_wardrobe_item", {
                item_type:this.wardrobe_item.item_type,
                brand_name:this.wardrobe_item.brand_name,
                size:this.wardrobe_item.size
            })
            if(result.status == 201) {
                this.$router.go(0)
            }
        }*/
    },

    async mounted() {
        let user = localStorage.getItem('user')

        if(!user) {
            this.$router.push({name:'SignUp'})
        }

        /*let result = await axios.get("http://0.0.0.0:3000/v1/it_wardrobe_item")
        let sizeType = await axios.get("http://0.0.0.0:3000/v1/it_garment_size_type")
        let sizes = await axios.get("http://0.0.0.0:3000/v1/it_garment_size_type_item")
        this.email = JSON.parse(user).email
        this.it_wardrobe_item = result.data
        this.it_garment_size_type = sizeType.data
        this.it_garment_size_type_item = sizes.data*/
    }
}
</script>

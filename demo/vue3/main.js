const { createApp, ref } = Vue

createApp({
    template: `
        <div v-if="show">
           hello {{msg}}
        </div>
    `,
    setup() {
        const show = ref(true)
        const msg = ref("gmpa")
        return {
            show,
            msg
        }
    }
}).mount('#root')

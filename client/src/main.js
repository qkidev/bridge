import { createApp } from 'vue';
import App from './App.vue'
import Vant from 'vant';
import {Toast} from 'vant';
import 'vant/lib/index.css';

const app = createApp(App);
app.use(Vant);
app.use(Toast)
app.mount('#app')
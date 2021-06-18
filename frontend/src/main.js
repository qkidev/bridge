import Vue from 'vue'
import App from './App.vue'
import '@/utils/rem'
import router from '@/router/router';
import './common/common.scss';
import { Toast } from 'vant';
import 'vant/lib/toast/style';
Vue.use(Toast);
import { Popup } from 'vant';
import 'vant/lib/popup/style';
Vue.use(Popup);
import { List } from 'vant';
import 'vant/lib/list/style';
Vue.use(List);

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
  data: {
    provider: null,
    signer: null,
  }
}).$mount('#app')

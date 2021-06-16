<!doctype html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Pair</title>
    <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css">
    <style>
        [v-cloak] {
            display: none;
        }
    </style>
</head>
<body>
<div id="app" v-cloak>
    <el-button type="primary" @click="show.create=true">添加跨链对</el-button>

    <el-dialog title="收货地址" :visible.sync="show.create">
        <el-form :model="form.create">
            <el-form-item label="活动名称">
                <el-input v-model="form.create.name"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button size="mini" @click="onCreate">提交</el-button>
            </el-form-item>
        </el-form>
    </el-dialog>
</div>

<script src="https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/ethers@5.3.1/dist/ethers.umd.min.js"></script>
<script src="https://unpkg.com/element-ui/lib/index.js"></script>
<script>
    const app = new Vue({
        el: '#app',
        data() {
            return {
                chains:@json($chains),
                tokens:@json($tokens),
                bridges:@json($bridges),
                show: {create: false},
                form: {create: {}}
            }
        },
        methods: {
            onCreate() {
                const url = this.form.create.chain.url
                const provider = new ethers.providers.JsonRpcProvider(url)
                if (this.checkPair(provider, this.form.create.toChain, this.form.create.toToken)) {
                    alert("跨链对已经存在")
                } else {

                }
            },
            async checkPair(provider, toChain, toToken) {
                try {
                    return await provider.tokens(toChain, toToken)
                } catch (e) {
                    return false
                }
            }
        }
    });
</script>
</body>
</html>

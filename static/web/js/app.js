const statusColors = {
    ok: 'olive',
    warn: 'orange',
    crit: 'red'
}

// box with overview information of a host
// currently not used!
Vue.component('host-box', {
    template: '#host-box',
    props: ['host'], // object with key, label, status

});

// Home dashboard (called by router)
const Home = {
    template: '#home',
    data() {
        // TODO get /api/list
        // debug:
        return {
            hosts: [
                {
                    key: 'linus',
                    label: 'Linus',
                    status: 'ok'
                },
                {
                    key: 'jarvis',
                    label: 'Jarvis',
                    status: 'warn'
                }
            ]
        }
    },
    methods: {
        statusLabel(host) {
            return '<a class="ui ' + statusColors[host.status] + ' small label">' + host.status + '</a>';
        },
        statusColor(host) {
            return statusColors[host.status];
        }
    }
}

// Host dashboard (called by router)
const Host = {
    template: '#host',
    props: ['host'], // host name
    data() {
        // TODO get /api/host/:host
    }
}

const router = new VueRouter({
    routes: [
        { path: '/', component: Home },
        { path: '/host/:host', component: Host }
    ]
})


var app = new Vue({
    el: '#app',
    router
})

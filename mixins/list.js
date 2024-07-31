import moment from "moment";

export default {
    data: () => ({
        pagination: {
            page: 1,
            pageSize: 10,
        },
    }),
    methods: {
        routerBack() {
            this.$router.push({name: 'admin-homepage'});
        },
        getData() {},
        moment(data, format = "DD/MM/YYYY HH:mm:ss") {
            if (moment(data).isValid()) {
                return moment(data).utc().add(7, "hours").format(format);
            }
            return data;
        },
        async handleTableChange(pagination) {
            this.pagination.page = pagination.current;
            await this.getData();
        },
    },
    created() {
        this.getData();
    },
};

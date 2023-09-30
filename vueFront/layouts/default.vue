<template>
    <v-card>
        <v-tabs v-model="tab" bg-color="primary">
            <v-tab to="/home">Home</v-tab>
            <v-tab to="/board">Board</v-tab>
            <userinfo />
        </v-tabs>

        <v-card-text>
            <v-window v-model="tab">
                <slot />
            </v-window>

        </v-card-text>
    </v-card>

</template>

<script>
export default {
    data: () => ({
        tab: null,
    }),
}
</script>
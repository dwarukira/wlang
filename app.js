const users = [
    {
        name: "Duncan",
        id: 1
    },
    {
        name: "Duncan @",
        id: 2
    },{
        name: "Duncan %",
        id: 3
    }
]

const newusers = users.filter((user ) => {

    return user.id === 3;
});

console.log(newusers);


application: {
    date.from : setUTCDate();
    date.to : setUTCDate();
    enforcement.value : number;
    enforcement.unit : Intl.NumberFormat.currency();
    activity.value : number;
    update(observation.commonwealth,
    date.from,
    date.to,
    enforcement.value,
    enforcement.unit,
    activity.value
    principles : {"Peoples are free and independent",
    "Peoples freedom and independence are to be respected by other Peoples",
    "Peoples are to observe treaties and undertakings",
    "Peoples are equal and are parties to the agreements that bind them",
    "Peoples are to observe a duty of non-intervention",
    "Peoples have the right of self-defense but no right to instigate war for reasons other than self-defense",
    "Peoples are to honour human rights",
    "Peoples are to observe certain specified restrictions in the conduct of war",
    "Peoples have a duty to assist other Peoples living under unfavourable conditions that prevent their having a just or decent political and social regime"});
},
observation.commonwealth contract: {
    update(observation.commonwealth):
        put(observation.commonwealth);
        return(observation.commonwealth);
}
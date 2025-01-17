application: {
    date.from : setUTCDate();
    date.to : setUTCDate();
    expenditure.value : number;
    expenditure.unit : Intl.NumberFormat.currency();
    activity.value : number;
    identity : Guid;
    owner : Organization;
    update(observation.service,
        date.from,
        date.to,
        expenditure.value,
        expenditure.unit,
        activity.value,
        identity,
        owner);
},
observation.service contract: {
    update(observation.service):
        put(observation.service);
        return(observation.service);
}
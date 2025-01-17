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
    activity.value);
},
observation.commonwealth contract: {
    update(observation.commonwealth):
        put(observation.commonwealth);
        return(observation.commonwealth);
},
observation.commonwealth interface: {
    Transactions : update;
    Endorsement Policy : Two Organization(MCOP);
}
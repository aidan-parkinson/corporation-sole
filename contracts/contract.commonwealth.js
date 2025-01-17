application: {
    peer : Organization(MCOP),
    client : Organization(COP),
    date.from : setUTCDate(),
    date.to : setUTCDate(),
    commonwealth.enforcement.value : number,
    commonwealth.enforcement.unit : Intl.NumberFormat.currency(),
    commonwealth.activity.value : number,
    update(observation.commonwealth, peer,
    client,
    date.from,
    date.to,
    commonwealth.enforcement.value,
    commonwealth.enforcement.unit,
    commonwealth.activity.value);
},
observation.commonwealth contract: {
    update(observation.commonwealth):
        put(observation.commonwealth);
        return(observation.commonwealth);
},
observation.commonwealth interface: {
    Transactions : update;
    Endorsement Policy : Any Organization(MCOP);
}
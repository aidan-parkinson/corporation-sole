application: {
    peer : Organization(MCOP);
    client : Organization(COP);
    date.from : setUTCDate();
    date.to : setUTCDate();
    service.expenditure.value : number;
    service.expenditure.unit : Intl.NumberFormat.currency();
    service.activity.value : number;
    service.identity : Guid;
    service.owner : Organization;
    update(observation.service.performance, peer,
        client,
        date.from,
        date.to,
        service.expenditure.value,
        service.expenditure.unit,
        service.activity.value,
        service.identity,
        service.owner);
},
observation.service contract: {
    update(observation.service):
        put(observation.service);
        return(observation.service);
},
observation.service.performance interface: {
    Transactions : update;
    Endorsement Policy : Any observation.service.owner(MCOP);
}
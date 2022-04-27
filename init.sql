INSERT INTO goflow.procdef (id, company_id, company_name, name, ver, resource, create_id, update_id, create_time, update_time, status) VALUES('1', '123', 'aaa', 'vacation apply', '1', '{
    "id": "process1",
    "start": {
        "name": "apply vacation",
        "id": "start",
        "formId" :"tt-task1-form",
        "nextId": "fork1"
    },
    "tasks": [
        {
            "name": "in 3 days,manager approve",
            "id": "task1",
            "nextId": "fork2",
            "prevId": "fork1",
            "formId" :"tt-task1-form",
            "approveType": 4
        },
        {
            "name": "more than 3 days, manager and ceo approve",
            "id": "task2",
            "nextId": "fork2",
            "prevId": "fork1",
            "formId" :"tt-task1-form",
            "actType": "and",
            "approveNum": 2,
            "approveType": 5,
            "approvers": [{
                "id": "3111",
                "name": "ceo tony"}
            ]
        }
    ],
    "forks": [
        {
            "name": "fork condition",
            "id": "fork1",
            "nextId": ["task1","task2"],
            "conds": "vdate<=3?0:1",
            "params": [{
                "name": "vdate",
                "type": "int"
            }]
        },{
            "name": "fork2 condition",
            "id": "fork2",
            "nextId": ["end1","end2"],
            "conds": "approve=='ok'?0:1",
            "params": [{
                "name": "approve",
                "type": "string"
            }]
        }
    ],
    "end": [{
        "name": "Approved End",
        "id": "end1"
    },{
        "name": "Rejected End",
        "id": "end2"
    }]
}', '1', '1', NULL, NULL, '0');

INSERT INTO goflow.form (id, form_name, form, form_desc, create_id, update_id, create_time, update_time, status) VALUES('tt-task1-form', 'vacation form', '<label>how many days you want vacation?</label><br>
        <input type="text" id="vdate" name="vdate"><br>
        <label>vacation reason ?</label><br>
        <input type="text" id="vreason" name="vreason"><br>', '{"vdate":"input","vreason":"input"}', NULL, NULL, NULL, NULL, NULL);


db.student.insert({
    name: "stud1",
    courses: [
        {
            name: "Programarea Calculatoarelor",
            year: 2020,
            series: "CB",
            abbreviation: "PC",
            labs: [
                {
                    name: "Laborator 1",
                    total: 0.25,
                    deadline: "05-19-2020",
                },
                {
                    name: "Laborator 2",
                    total: 0.25,
                    deadline: "05-26-2020",
                }
            ],
            homeworks: []
        }
    ]
});

db.student.insert({
    name: "stud4",
    courses: [
        {
            name: "Algoritmi Paraleli si Distribuiti",
            year: 2020,
            series: "CC",
            abbreviation: "APD",
            labs: [
                {
                    name: "Laborator 1",
                    total: 0.25,
                    deadline: "05-19-2020",
                },
                {
                    name: "Laborator 2",
                    total: 0.25,
                    deadline: "05-26-2020",
                }
            ],
            homeworks: [
                {
                    name: "Tema 1",
                    total: 1,
                    deadline: "05-30-2020",
                }
            ]
        },
        {
            name: "Programarea Calculatoarelor",
            year: 2020,
            series: "CC",
            abbreviation: "PC",
            labs: [
                {
                    name: "Laborator 1",
                    result: 80,
                    total: 0.25,
                    deadline: "10-16-2018",
                }
            ],
            homeworks: [
                {
                    name: "Tema 1",
                    result: 90,
                    total: 1,
                    deadline: "11-09-2018",
                }
            ]
        }
    ]
});

db.assistant.insert({
    name: "stud4",
    courses: [
        {
            name: "Programarea Calculatoarelor",
            year: 2020,
            series: "CB",
            abbreviation: "PC",
            assignments: [
                {
                    name: "Laborator 1",
                    students: [
                        {
                            name: "stud1",
                            grade: "",
                            gradetime: "",
                            graded_by: ""
                        },
                        {
                            name: "stud2",
                            grade: "",
                            gradetime: "",
                            graded_by: ""
                        },
                        {
                            name: "stud3",
                            grade: "",
                            gradetime: "",
                            graded_by: ""
                        }
                    ]
                }
            ]
        }
    ]
});
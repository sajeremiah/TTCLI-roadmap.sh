# CLI Task Tracker

## Project task: https://roadmap.sh/projects/task-tracker

### How to run

#### Clone the repository and run the following command:

```
git clone https://github.com/sajeremiah/TTCLI-roadmap.sh
cd ./TTCLI-roadmap.sh
```

#### Run the following command to build and run the project:

```sh
make build                      > build in ./bin
make run <ARGUMENTS>            > build & run with args in cmd line
```

## To add a task
```sh
./bin/task-cli add "Buy groceries"
```

## To update a task
```sh
./bin/task-cli update 1 "Buy groceries and cook dinner"
```
## To delete a task
```sh
./bin/task-cli delete 1
```
## To mark a task as in progress/done/todo
```sh
./bin/task-cli mark-in-progress 1
./bin/task-cli mark-done 1
```
## To list all tasks
```sh
./bin/task-cli list
./bin/task-cli list done
./bin/task-cli list in-progress
```

## Json-file example view
```json
{
   "nextId": 3,
   "tasks": [
      {
         "id": 1,
         "description": "Buy groceries",
         "status": "in-progress",
         "created_at": "2026-09-21T11:45:01.314242589+03:00",
         "updated_at": null
      },
      {
         "id": 2,
         "description": "Complete Task Tracker proj",
         "status": "done",
         "created_at": "2026-09-21T11:45:38.242574899+03:00",
         "updated_at": "2026-09-21T11:46:55.324035152+03:00"
      }
   ]
}
```
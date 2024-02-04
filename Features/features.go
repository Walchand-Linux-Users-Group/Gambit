package features

type Project struct {
	Project_Name      string
	Assigned_to       string
	Completion_Status bool
	Cost              int
}

type Manager struct {
	Projects []Project
}

func NewClass() *Manager {
	return &Manager{
		Projects: []Project{},
	}
}

func (M Manager) Hello() string {
	return "Hello From Manager"
}

func (M *Manager) Insert(cost int, status bool, assigned_to, project_name string) {

	// if cost < 0 || assigned_to == "" || project_name == "" {
	// 	fmt.Println("Please Provide Appropriate Values")
	// 	return
	// }

	// temp := Project{
	// 	Project_Name:      project_name,
	// 	Assigned_to:       assigned_to,
	// 	Completion_Status: status,
	// 	Cost:              cost,
	// }

	// M.Projects = append(M.Projects, temp)
}

func (M *Manager) Sum() {
	// sum := 0
	// for _, v := range M.Projects {
	// 	sum += v.Cost
	// }
	// fmt.Println(sum)
}

func (M *Manager) List() {
	// fmt.Println("----- List Of Projects ------")
	// for i, v := range M.Projects {
	// 	fmt.Printf("\nIndex: %v\nProject Name: %v\nAssigned To: %v\nCompletion Status: %v\nCost: %v\n", i, v.Project_Name, v.Assigned_to, v.Completion_Status, v.Cost)
	// }
}

func (M *Manager) Update(status bool, assigned_to string, project_name string, index, cost int) {
	// if cost < 0 || assigned_to == "" || project_name == "" {
	// 	fmt.Println("Please Provide Appropriate Values")
	// 	return
	// }
	// for i := range M.Projects {
	// 	if i == index {
	// 		M.Projects[i] = Project{
	// 			Project_Name:      project_name,
	// 			Assigned_to:       assigned_to,
	// 			Cost:              cost,
	// 			Completion_Status: status,
	// 		}
	// 		break
	// 	}
	// }
}

func (M *Manager) Delete(index int) {
	// for i := range M.Projects {
	// 	if i == index {
	// 		M.Projects = append(M.Projects[:index], M.Projects[index+1:]...)
	// 		return
	// 	}
	// }
	// fmt.Println("Index out of bound")
}

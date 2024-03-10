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




// ------------  CHANGE FOLLOWING CODE INORDER TO COMPLETE THE TASKS ------------------- //

func (M Manager) Hello() string {

	// UNCOMMENT THE FOLLOWING CODE INORDER TO GET THE INITIAL POINTS
	
	// return "Hello From Manager"
}

func (M *Manager) Insert(cost int, status bool, assigned_to, project_name string) {
	
}

func (M *Manager) Sum() {
	
}

func (M *Manager) List() {

	// CODE OF THE 'List' FUNCTION IS GIVEN BELOW, UNCOMMENT AND USE IT DIRECTLY :)
	
	// fmt.Println("----- List Of Projects ------")
	// for i, v := range M.Projects {
	// 	fmt.Printf("\nIndex: %v\nProject Name: %v\nAssigned To: %v\nCompletion Status: %v\nCost: %v\n", i, v.Project_Name, v.Assigned_to, v.Completion_Status, v.Cost)
	// }
}

func (M *Manager) Update(status bool, assigned_to string, project_name string, index, cost int) {
	
}

func (M *Manager) Delete(index int) {
	
}

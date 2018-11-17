package models

type RoleType string

// Don't change the names since these values are usually saved in the
// database
const (
	RoleTypeUndefined   RoleType = "undefined"
	RoleTypeNormal      RoleType = "normal"
	RoleTypeCircle      RoleType = "circle"
	RoleTypeLeadLink    RoleType = "leadlink"
	RoleTypeRepLink     RoleType = "replink"
	RoleTypeFacilitator RoleType = "facilitator"
	RoleTypeEngager     RoleType = "engager"
	RoleTypeChampion    RoleType = "champion"
	RoleTypeScout       RoleType = "scout"
	RoleTypeMagister    RoleType = "magister"
	RoleTypeMangler     RoleType = "mangler"
	RoleTypeSecretary   RoleType = "secretary"
)

func (r RoleType) IsCoreRoleType() bool {
	return r == RoleTypeLeadLink ||
		r == RoleTypeRepLink ||
		r == RoleTypeFacilitator ||
		r == RoleTypeEngager ||
		r == RoleTypeChampion ||
		r == RoleTypeScout ||
		r == RoleTypeMagister ||
		r == RoleTypeMangler ||
		r == RoleTypeSecretary
}

func (r RoleType) String() string {
	return string(r)
}

func RoleTypeFromString(r string) RoleType {
	switch r {
	case "normal":
		return RoleTypeNormal
	case "circle":
		return RoleTypeCircle
	case "leadlink":
		return RoleTypeLeadLink
	case "replink":
		return RoleTypeRepLink
	case "facilitator":
		return RoleTypeFacilitator
	case "engager":
		return RoleTypeEngager
	case "champion":
		return RoleTypeChampion
	case "scout":
		return RoleTypeScout
	case "magister":
		return RoleTypeMagister
	case "mangler":
		return RoleTypeMangler
	case "secretary":
		return RoleTypeSecretary
	default:
		return RoleTypeUndefined
	}
}

type Role struct {
	Vertex
	RoleType RoleType
	Depth    int32
	Name     string
	Purpose  string
}

type Roles []*Role

func (r Roles) Len() int           { return len(r) }
func (r Roles) Less(i, j int) bool { return r[i].Name < r[j].Name }
func (r Roles) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

type RoleAdditionalContent struct {
	Vertex
	Content string
}

type MemberCirclePermissions struct {
	AssignChildCircleLeadLink   bool
	AssignChildRoleMembers      bool
	ManageChildRoles            bool
	AssignCircleDirectMembers   bool
	AssignCircleCoreRoles       bool
	ManageRoleAdditionalContent bool
	// special cases for root circle
	AssignRootCircleLeadLink bool
	ManageRootCircle         bool
}

type CoreRoleDefinition struct {
	Role             *Role
	Domains          []*Domain
	Accountabilities []*Accountability
}

func GetCoreRoles() []*CoreRoleDefinition {
	return []*CoreRoleDefinition{
		{
			Role: &Role{
				Name:     "Sircle Leader",
				RoleType: RoleTypeLeadLink,
				Purpose:  "Sircle Leader holds the Purpose of the overall Sircle",
			},
			Domains: []*Domain{
				{Description: "Allocate Role within Sircle"},
			},
			Accountabilities: []*Accountability{
				{Description: "Structure Sircle’s Governance to enact its Purpose and Accountabilities"},
				{Description: "Sircle's Act Owner"},
				{Description: "Allocating the Circle's resources across its various Projects and/or Roles"},
				{Description: "Assign people to Sircle's Roles and Monitor the fit"},
				{Description: "Offer feedback to enhance fit and re-assign Roles to other people when it could be useful for enhancing fit"},
				{Description: "Allocate Sircle's resources through its various Projects and/or Roles"},
				{Description: "Establish priorities and Strategies for the Sircle"},
				{Description: "Define a more general Strategy for the Sircle, or multiple Strategies, which are heuristics that guide the Sircle's Roles in self-identifying priorities on an ongoing basis"},
				{Description: "Define kpi for the Sircle"},
				{Description: "Remove constraints within the Sircle through Super-Sircle's Purpose and Accountabilities"},
				{Description: "Remove impediments (Facilitator)"},
				{Description: "Team coaching"},
				{Description: "Write and keep up to date Value Proposition (VP) document"},
				{Description: "Write a blog + twitter/linkedin/wetalk updates about new trends and events and nurture followers"},
				{Description: "Support Hr caring people through leveraging Sorint values"},
			},
		},
		{
			Role: &Role{
				Name:     "Rep Link",
				RoleType: RoleTypeRepLink,
				Purpose:  "Within the Super-Circle, the Rep Link holds the Purpose of the SubCircle; within the Sub-Circle, the Rep Link’s Purpose is: Tensions relevant to process in the Super-Circle channeled out and resolved",
			},
			Accountabilities: []*Accountability{
				{Description: "Removing constraints within the broader Organization that limit the Sub-Circle"},
				{Description: "Seeking to understand Tensions conveyed by Sub-Circle Circle Members, and discerning those appropriate to process in the Super-Circle"},
				{Description: "Providing visibility to the Super-Circle into the health of the Sub-Circle, including reporting on any metrics or checklist items assigned to the whole Sub-Circle"},
			},
		},
		{
			Role: &Role{
				Name:     "Facilitator",
				RoleType: RoleTypeFacilitator,
				Purpose:  "Circle governance and operational practices aligned with the Constitution",
			},
			Accountabilities: []*Accountability{
				{Description: "Facilitating the Circle’s constitutionally-required meetings"},
				{Description: "Auditing the meetings and records of Sub-Circles as needed, and declaring a Process Breakdown upon discovering a pattern of behavior that conflicts with the rules of the Constitution"},
			},
		},
		{
			Role: &Role{
				Name:     "Engager",
				RoleType: RoleTypeEngager,
				Purpose:  "Extreme customer satisfaction and proposing real added value solutions with visible Sorint signature",
			},
			Domains: []*Domain{
				{Description: "Technical Sircles"},
			},
			Accountabilities: []*Accountability{
				{Description: "Define the number of resources needed to deliver the activity and their required skills for each one"},
				{Description: "Provide the Statement of work for the activity which allows Sorint to build a trade offer defining  solution studies (proposal, solution design, etc ...)"},
				{Description: "Build Use Cases and solutions description to be used on sales activities"},
				{Description: "Attend meetings when Business Developer asks for support for his/her sales activities"},
				{Description: "Provide Product Service Presentation"},
				{Description: "Build and run Poc"},
				{Description: "Show Different Technologies through Use Cases"},
				{Description: "Show Product's Best Practices"},
			},
		},
		{
			Role: &Role{
				Name:     "Champion",
				RoleType: RoleTypeChampion,
				Purpose:  "Extreme customer satisfaction and establishing valued relationship with visible Sorint signature, the Sorint ambassador",
			},
			Domains: []*Domain{
				{Description: "Customer Sircle"},
			},
			Accountabilities: []*Accountability{
				{Description: "Identify and map customer organization and various managers"},
				{Description: "Study the technical sircles's value proposition"},
				{Description: "Identify successful cases which could be offered to the client"},
				{Description: "Arrange regular meeting with the customer's referrals in order to share the vp"},
				{Description: "Open opportunities for any intercepted customer needs"},
				{Description: "Align with and compare to the Sircle leader for opportunities"},
				{Description: "Report competitors presence in the site (name and where)"},
				{Description: "Report negative and positive customer feedback on our work"},
				{Description: "Report negative and positive customer feedback on competitors works"},
				{Description: "Report any changes in the site that could lead to opportunities, instability or problems"},
				{Description: "Report intentions in order to purchase new products"},
				{Description: "Be aware of all customer-related opportunities"},
				{Description: "Organize Meetings/phone calls with Sorint Business Developer for alignment"},
			},
		},		
		{
			Role: &Role{
				Name:     "Talent Handler",
				RoleType: RoleTypeScout,
				Purpose:  "Support HR in managing new placements",
			},
			Domains: []*Domain{
				{Description: "Customer Sircle"},
			},
			Accountabilities: []*Accountability{
				{Description: "Check CVs offers and giving feedback to HR and candidates"},
				{Description: "Conduct Job pre-interview/interview with candidate"},
				{Description: "Read and edit CVs in order to the manifest required features"},
				{Description: "Prepare a candidate for the customer interview"},
				{Description: "Submit a candidate for the customer interview"},
				{Description: "Support HR in technical talks with candidates"},
			},
		},
		{
			Role: &Role{
				Name:     "Academier",
				RoleType: RoleTypeMagister,
				Purpose:  "Circle governance and operational practices aligned with the Constitution",
			},
			Domains: []*Domain{
				{Description: "Sircles Core Members"},
			},
			Accountabilities: []*Accountability{
				{Description: "Explicit the core member education needs in order to achieve the Sircle purpose"},
				{Description: "Write a Training Path"},
			},
		},
		{
			Role: &Role{
				Name:     "Planner",
				RoleType: RoleTypeMangler,
				Purpose:  "After the Sircle leader assign member to a Role, The planner takes care of the right ACO, ACI activities assignation and scheduling for all the members",
			},
			Domains: []*Domain{
				{Description: "Services Activities under Sircle's responsibility"},
			},
			Accountabilities: []*Accountability{
				{Description: "Keep update the schedule"},
				{Description: "Schedule properly to be able to complete Projects/services successfully"},
				{Description: "Keep update the consultants and project leads on a regular basis"},
				{Description: "Identify bottlenecks and priorities conflicts reporting to Sircle Leader and PMO"},
				{Description: "Manage resources requests conflicts"},
				{Description: "Inform PMO about People Availability"},
			},
		},
		{
			Role: &Role{
				Name:     "Secretary",
				RoleType: RoleTypeSecretary,
				Purpose:  "Steward and stabilize the Circle’s formal records and record-keeping process",
			},
			Domains: []*Domain{
				{Description: "All constitutionally-required records of the Circle"},
			},
			Accountabilities: []*Accountability{
				{Description: "Scheduling the Circle’s required meetings, and notifying all Core Circle Members of scheduled times and locations"},
				{Description: "Capturing and publishing the outputs of the Circle’s required meetings, and maintaining a compiled view of the Circle’s current Governance, checklist items, and metrics"},
				{Description: "Interpreting Governance and the Constitution upon request"},
			},
		},
	}
}

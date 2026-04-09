// Integration registry — add new providers here to expose them in the
// setup modal. Each entry declares the credential fields it needs and a
// `fetchResources` function returning the items shown in the final step.
//
// Backend wiring is intentionally left out for now: `fetchResources`
// returns mock data so the UI can be built end-to-end.

export const integrations = [
  {
    id: "azure-devops",
    label: "Azure DevOps",
    credentialFields: [
      {
        key: "pat",
        label: "Personal Access Token",
        type: "password",
        placeholder: "Enter your Azure DevOps PAT",
        help: "Generate one from User Settings → Personal Access Tokens.",
      },
    ],
    resourceLabel: "Project",
    async fetchResources(/* credentials */) {
      return [
        { id: "ado-1", name: "Contoso Web" },
        { id: "ado-2", name: "Contoso Mobile" },
        { id: "ado-3", name: "Internal Tools" },
        { id: "ado-4", name: "Data Platform" },
      ];
    },
  },
  {
    id: "github",
    label: "GitHub",
    credentialFields: [],
    resourceLabel: "Repository",
    async fetchResources() {
      return [];
    },
  },
  {
    id: "jira",
    label: "Jira",
    credentialFields: [],
    resourceLabel: "Board",
    async fetchResources() {
      return [];
    },
  },
];

export function getIntegration(id) {
  return integrations.find((i) => i.id === id);
}

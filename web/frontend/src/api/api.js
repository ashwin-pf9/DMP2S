import API_BASE_URL from "../config";

const getAuthToken = () => {
    return localStorage.getItem("token"); // Retrieve token from localStorage
  };
  
  export const fetchPipelines = async () => {
    const response = await fetch(`${API_BASE_URL}/pipelines`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${getAuthToken()}`,
        "Content-Type": "application/json"
      }
    });
  
    if (!response.ok) throw new Error("Failed to fetch pipelines");
    return response.json();
  };

export const fetchStages = async (pipelineId) => {
  const response = await fetch(`${API_BASE_URL}/pipelines/${pipelineId}/stages`,{
    method: "GET",
    headers: {
      "Authorization": `Bearer ${getAuthToken()}`,
      "Content-Type": "application/json"
    }
  });
  if (!response.ok) throw new Error("Failed to fetch stages");
  return response.json();
};

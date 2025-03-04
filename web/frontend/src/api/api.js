import API_BASE_URL from "../config";

const getAuthToken = () => {
    return localStorage.getItem("token"); // Retrieve token from localStorage
  };
  
  export const fetchPipelines = async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/pipelines`, {
        method: "GET",
        headers: {
          "Authorization": `Bearer ${getAuthToken()}`,
          "Content-Type": "application/json"
        }
      });
  
      if (!response.ok) {
        throw new Error(`Failed to fetch pipelines: ${response.statusText}`);
      }
  
      const data = await response.json();
      console.log("Fetched Pipelines:", data); // Debugging output
  
      if (!data || !Array.isArray(data)) {
        console.error("Invalid pipeline data format:", data);
        throw new Error("Invalid pipeline data format");
      }
  
      return data;
    } catch (error) {
      console.error("Error fetching pipelines:", error);
      return [];
    }
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

export const createPipeline = async (pipelineName) => {
  try {
    const response = await fetch(`${API_BASE_URL}/pipelines/create`, {
      method: "POST",
      headers: {
        "Authorization": `Bearer ${getAuthToken()}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name: pipelineName }),
    });

    if (!response.ok) {
      throw new Error(`Failed to create pipeline: ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Error creating pipeline:", error);
    throw error;
  }
};
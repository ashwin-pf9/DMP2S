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

    if (!data || !Array.isArray(data)) {
      console.error("Invalid pipeline data format:", data);
      throw new Error("Invalid pipeline data format");
    }

    return data;
  } catch (error) {
    if (error.message === "Invalid authorization token") {
      alert("Session expired. Redirecting to login...");
      window.location.href = "/login"; // Redirect user to login page
    } else if (error.message.includes("Failed to fetch") || error.message.includes("NetworkError")) {
      console.error("Backend is unreachable:", error);
      throw new Error("Backend is unreachable. Please try again later.");
    }


    console.error("Error fetching pipelines:", error);
    throw error;  // Throw error instead of returning an empty array
  }
};


export const fetchStages = async (pipelineId) => {
  try {
    const response = await fetch(`${API_BASE_URL}/pipelines/${pipelineId}/stages`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${getAuthToken()}`,
        "Content-Type": "application/json"
      }
    });
    if (!response.ok) {
      if (response.status === 401) {
        throw new Error("Unauthorized");
      }
      throw new Error("Failed to fetch stages");
    }
    return response.json();
  } catch (error) {
    console.error("Error fetching stages:", error);
  }


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
      if (response.status === 401) {
        throw new Error("Unauthorized");
      }
      throw new Error(`Failed to create pipeline: ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    alert(`Error: ${error.message}`);
  }
};

export const deletePipeline = async (pipelineId) => {
  try {
    const response = await fetch(`${API_BASE_URL}/pipelines/${pipelineId}/delete`, {
      method: "POST",
      headers: {
        "Authorization": `Bearer ${getAuthToken()}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      if (response.status === 401) {
        console.log("log - Unauthorized")
        throw new Error("Unauthorized");
      }
      throw new Error(`Failed to delete pipeline: ${response.statusText}`);
    }

    alert("Pipeline deleted successfully!");
    window.location.href = "/dashboard";

    return await response.json();
  } catch (error) {
    if (error.message === "Unauthorized") {
      alert("Session expired. Redirecting to login...");
      window.location.href = "/login"; // Redirect user to login page
    } else {
      alert(`Error: ${error.message}`);
    }
  }
};

export const startPipeline = async (pipelineId) => {
  try {
    const response = await fetch(`${API_BASE_URL}/pipelines/${pipelineId}/start`, {
      method: "POST",
      headers: {
        "Authorization": `Bearer ${getAuthToken()}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      if (response.status === 401) {
        console.log("log - Unauthorized")
        throw new Error("Unauthorized");
      }
      throw new Error(`Failed to start pipeline: ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    if (error.message === "Unauthorized") {
      alert("Session expired. Redirecting to login...");
      window.location.href = "/login"; // Redirect user to login page
    } else {
      alert(`Error: ${error.message}`);
    }
  }
};

export const createStage = async (pipelineId, stage) => {
  try {
    console.log(pipelineId)
    console.log(stage)
    const response = await fetch(`${API_BASE_URL}/pipelines/${pipelineId}/stages/add`, {
      method: "POST",
      headers: {
        "Authorization": `Bearer ${getAuthToken()}`,
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        name: stage.name,
        pipeline_id: pipelineId
      })
    });


    if (!response.ok) {
      throw new Error(`Failed to create stage: ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    alert(`Error: ${error.message}`);
  }
};

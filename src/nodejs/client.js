const axios = require('axios');

class Client {
    constructor(baseURL) {
        this.client = axios.create({
            baseURL: baseURL
        });
    }

    async makeRequest(endpoint) {
        try {
            const response = await this.client.get(endpoint);
            return response.data;
        } catch (error) {
            console.error('Error making request:', error);
            throw error;
        }
    }
}

module.exports = Client;
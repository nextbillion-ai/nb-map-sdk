class SDKError extends Error {
    constructor(message) {
        super(message);
        this.name = "SDKError";
        
    }
}

module.exports = {
    SDKError,
    // add more errors
};



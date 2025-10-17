const express = require('express');
const axios = require('axios');

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(express.json());

// Health check endpoint
app.get('/health', (req, res) => {
  res.json({
    status: 'healthy',
    timestamp: new Date().toISOString(),
    service: 'api-gateway',
    version: '1.0.0'
  });
});

// Reverse proxy endpoint
app.get('/api/proxy', async (req, res) => {
  try {
    const targetUrl = req.query.url;
    
    if (!targetUrl) {
      return res.status(400).json({
        error: 'Missing url parameter',
        message: 'Please provide a target URL via ?url=<target_url>'
      });
    }

    // Forward request to target URL
    const response = await axios.get(targetUrl);
    
    // Return proxied response
    res.json(response.data);
  } catch (error) {
    res.status(error.response?.status || 500).json({
      error: 'Proxy request failed',
      message: error.message
    });
  }
});

// Root endpoint
app.get('/', (req, res) => {
  res.json({
    service: 'API Gateway Service',
    version: '1.0.0',
    endpoints: {
      health: '/health',
      proxy: '/api/proxy?url=<target_url>'
    }
  });
});

// Start server
app.listen(PORT, () => {
  console.log(`🚀 API Gateway running on http://localhost:${PORT}`);
  console.log(`📊 Health check: http://localhost:${PORT}/health`);
  console.log(`🔄 Proxy endpoint: http://localhost:${PORT}/api/proxy?url=<target_url>`);
});

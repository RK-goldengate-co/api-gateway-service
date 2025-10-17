from fastapi import FastAPI, Query, HTTPException
from datetime import datetime
import httpx
import uvicorn

app = FastAPI(
    title="API Gateway Service",
    description="API Gateway với reverse proxy và health check",
    version="1.0.0"
)

@app.get("/")
async def root():
    """Root endpoint"""
    return {
        "service": "API Gateway Service",
        "version": "1.0.0",
        "endpoints": {
            "health": "/health",
            "proxy": "/api/proxy?url=<target_url>",
            "docs": "/docs"
        }
    }

@app.get("/health")
async def health_check():
    """Health check endpoint"""
    return {
        "status": "healthy",
        "timestamp": datetime.utcnow().isoformat() + "Z",
        "service": "api-gateway",
        "version": "1.0.0"
    }

@app.get("/api/proxy")
async def proxy_request(url: str = Query(..., description="Target URL to proxy")):
    """Reverse proxy endpoint"""
    try:
        async with httpx.AsyncClient() as client:
            response = await client.get(url)
            return response.json()
    except httpx.RequestError as e:
        raise HTTPException(
            status_code=500,
            detail={
                "error": "Proxy request failed",
                "message": str(e)
            }
        )
    except Exception as e:
        raise HTTPException(
            status_code=500,
            detail={
                "error": "Internal server error",
                "message": str(e)
            }
        )

if __name__ == "__main__":
    print("🚀 API Gateway running on http://localhost:8000")
    print("📊 Health check: http://localhost:8000/health")
    print("🔄 Proxy endpoint: http://localhost:8000/api/proxy?url=<target_url>")
    print("📚 API Documentation: http://localhost:8000/docs")
    uvicorn.run(app, host="0.0.0.0", port=8000)

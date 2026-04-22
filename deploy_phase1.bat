@echo off
SETLOCAL EnableDelayedExpansion

echo ============================================================
echo [729 HOLDINGS] RAAJJE HEADLINES - PHASE 1 DEPLOYMENT
echo ============================================================
echo.

:: 1. Build and Deploy Web Portal
echo [STEP 1/4] INITIALIZING WEB PORTAL BUILD...
cd raajje-portal
call npm install --silent
call npm run build
if %ERRORLEVEL% NEQ 0 (
    echo Error during build. Aborting.
    exit /b %ERRORLEVEL%
)
cd ..

echo [STEP 2/4] SYNCING TO ADHU.SPACE (FIREBASE)...
:: Only deploy hosting and related config
call npx firebase deploy --only hosting,dataconnect,firestore,auth
if %ERRORLEVEL% NEQ 0 (
    echo Firebase deployment failed.
)

:: 2. Deploy Edge Node / Aggregator to Cloud Run
echo [STEP 3/4] DEPLOYING AGENTIC EDGE NODE (CLOUD RUN)...
:: Use the cloudbuild.yaml to automate the build/push/deploy
call gcloud builds submit --config cloudbuild.yaml . --ignore-file=exclude.txt
if %ERRORLEVEL% NEQ 0 (
    echo Cloud Build failed. Checks logs for details.
)

:: 3. Prepare Android Assets
echo [STEP 4/4] GENERATING ANDROID RELEASE ASSETS (AAB)...
cd android_executive
call gradlew.bat bundleRelease
if %ERRORLEVEL% NEQ 0 (
    echo Android build failed.
)
cd ..

echo.
echo ============================================================
echo DEPLOYMENT PHASE 1 COMPLETE - ECOSYSTEM "READY FOR TAKEOFF"
echo ============================================================
pause

# read the workflow template
WORKFLOW_SERVICE_TEMPLATE=$(cat .github/templates/service-template.yml)

# iterate each route in routes directory
for SERVICE_NAME in $(ls servers); do
    # replace template service placeholder with service name
    echo "Generating workflow for service ${SERVICE_NAME}"

    # replace template route placeholder with route name
    WORKFLOW=$(echo "${WORKFLOW_SERVICE_TEMPLATE}" | sed "s/{{SERVICE_NAME}}/${SERVICE_NAME}/g")

    # save workflow to .github/workflows/{ROUTE}
    echo "${WORKFLOW}" > .github/workflows/${SERVICE_NAME}-service.yml
done
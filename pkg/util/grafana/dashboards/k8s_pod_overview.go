package grafana

const PodsOverviewKubePi = `{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": {
          "type": "datasource",
          "uid": "grafana"
        },
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "target": {
          "limit": 100,
          "matchAny": false,
          "tags": [],
          "type": "dashboard"
        },
        "type": "dashboard"
      }
    ]
  },
  "description": "This is a modern 'Pods View' dashboard for your Kubernetes cluster(s). Made for kube-prometheus-stack and take advantage of the latest Grafana features. GitHub repository: https://github.com/dotdc/grafana-dashboards-kubernetes",
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 1,
  "links": [],
  "liveNow": false,
  "panels": [
    {
      "collapsed": false,
      "datasource": {
        "type": "datasource",
        "uid": "grafana"
      },
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 0
      },
      "id": 43,
      "panels": [],
      "targets": [
        {
          "datasource": {
            "type": "datasource",
            "uid": "grafana"
          },
          "refId": "A"
        }
      ],
      "title": "Information",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "description": "Panel only works when a single pod is selected.",
      "fieldConfig": {
        "defaults": {
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "rgb(255, 255, 255)",
                "value": null
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 2,
        "w": 12,
        "x": 0,
        "y": 1
      },
      "id": 2,
      "options": {
        "colorMode": "none",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "percentChangeColorMode": "standard",
        "reduceOptions": {
          "calcs": [
            "mean"
          ],
          "fields": "",
          "values": false
        },
        "showPercentChange": false,
        "textMode": "name",
        "wideLayout": true
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "kube_pod_info{namespace=\"$namespace\", pod=\"$pod\", cluster=\"$cluster\"}",
          "instant": true,
          "interval": "",
          "legendFormat": "{{ created_by_kind }}: {{ created_by_name }}",
          "refId": "A"
        }
      ],
      "title": "Created by",
      "type": "stat"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "description": "Panel only works when a single pod is selected.",
      "fieldConfig": {
        "defaults": {
          "links": [
            {
              "title": "",
              "url": "/d/k8s_views_nodes/kubernetes-views-nodes?var-datasource=${datasource}&var-node=${__field.labels.node}"
            }
          ],
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "rgb(255, 255, 255)",
                "value": null
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 2,
        "w": 6,
        "x": 12,
        "y": 1
      },
      "id": 33,
      "options": {
        "colorMode": "none",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "percentChangeColorMode": "standard",
        "reduceOptions": {
          "calcs": [
            "mean"
          ],
          "fields": "",
          "values": false
        },
        "showPercentChange": false,
        "textMode": "name",
        "wideLayout": true
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "kube_pod_info{namespace=\"$namespace\", pod=\"$pod\", cluster=\"$cluster\"}",
          "instant": true,
          "interval": "",
          "legendFormat": "{{ node }}",
          "refId": "A"
        }
      ],
      "title": "Running on",
      "type": "stat"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "description": "Panel only works when a single pod is selected.",
      "fieldConfig": {
        "defaults": {
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "rgb(255, 255, 255)",
                "value": null
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 2,
        "w": 6,
        "x": 18,
        "y": 1
      },
      "id": 41,
      "options": {
        "colorMode": "none",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "percentChangeColorMode": "standard",
        "reduceOptions": {
          "calcs": [
            "mean"
          ],
          "fields": "",
          "values": false
        },
        "showPercentChange": false,
        "textMode": "name",
        "wideLayout": true
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "kube_pod_info{namespace=\"$namespace\", pod=\"$pod\", cluster=\"$cluster\"}",
          "instant": true,
          "interval": "",
          "legendFormat": "{{ pod_ip }}",
          "refId": "A"
        }
      ],
      "title": "Pod IP",
      "type": "stat"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "description": "Panel only works when a single pod is selected.",
      "fieldConfig": {
        "defaults": {
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "rgb(255, 255, 255)",
                "value": null
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 2,
        "w": 5,
        "x": 0,
        "y": 3
      },
      "id": 52,
      "options": {
        "colorMode": "none",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "percentChangeColorMode": "standard",
        "reduceOptions": {
          "calcs": [
            "mean"
          ],
          "fields": "",
          "values": false
        },
        "showPercentChange": false,
        "textMode": "name",
        "wideLayout": true
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "kube_pod_info{namespace=\"$namespace\", pod=\"$pod\", priority_class!=\"\", cluster=\"$cluster\"}",
          "format": "time_series",
          "instant": true,
          "interval": "",
          "legendFormat": "{{ priority_class }}",
          "range": false,
          "refId": "A"
        }
      ],
      "title": "Priority Class",
      "type": "stat"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "description": "Panel only works when a single pod is selected.",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              }
            ]
          },
          "unit": "none"
        },
        "overrides": [
          {
            "matcher": {
              "id": "byName",
              "options": "Burstable"
            },
            "properties": [
              {
                "id": "color",
                "value": {
                  "fixedColor": "red",
                  "mode": "fixed"
                }
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "BestEffort"
            },
            "properties": [
              {
                "id": "color",
                "value": {
                  "fixedColor": "orange",
                  "mode": "fixed"
                }
              }
            ]
          }
        ]
      },
      "gridPos": {
        "h": 2,
        "w": 7,
        "x": 5,
        "y": 3
      },
      "id": 53,
      "options": {
        "colorMode": "value",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "percentChangeColorMode": "standard",
        "reduceOptions": {
          "calcs": [],
          "fields": "",
          "values": false
        },
        "showPercentChange": false,
        "textMode": "name",
        "wideLayout": true
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "kube_pod_status_qos_class{namespace=\"$namespace\", pod=\"$pod\", cluster=\"$cluster\"} > 0",
          "instant": true,
          "interval": "",
          "legendFormat": "{{ qos_class }}",
          "refId": "A"
        }
      ],
      "title": "QOS Class",
      "type": "stat"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "description": "Panel only works when a single pod is selected.",
      "fieldConfig": {
        "defaults": {
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "text",
                "value": null
              },
              {
                "color": "red",
                "value": 1
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 2,
        "w": 6,
        "x": 12,
        "y": 3
      },
      "id": 56,
      "options": {
        "colorMode": "value",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "percentChangeColorMode": "standard",
        "reduceOptions": {
          "calcs": [],
          "fields": "",
          "values": false
        },
        "showPercentChange": false,
        "textMode": "name",
        "wideLayout": true
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "kube_pod_container_status_last_terminated_reason{namespace=\"$namespace\", pod=\"$pod\", cluster=\"$cluster\"}",
          "instant": true,
          "interval": "",
          "legendFormat": "{{ reason }}",
          "refId": "A"
        }
      ],
      "title": "Last Terminated Reason",
      "type": "stat"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "description": "Panel only works when a single pod is selected.",
      "fieldConfig": {
        "defaults": {
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "text",
                "value": null
              },
              {
                "color": "red",
                "value": 1
              },
              {
                "color": "#EAB839",
                "value": 2
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 2,
        "w": 6,
        "x": 18,
        "y": 3
      },
      "id": 57,
      "options": {
        "colorMode": "value",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "percentChangeColorMode": "standard",
        "reduceOptions": {
          "calcs": [],
          "fields": "",
          "values": true
        },
        "showPercentChange": false,
        "textMode": "value",
        "wideLayout": true
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "kube_pod_container_status_last_terminated_exitcode{namespace=\"$namespace\", pod=\"$pod\", cluster=\"$cluster\"}",
          "instant": true,
          "interval": "",
          "legendFormat": "__auto",
          "range": false,
          "refId": "A"
        }
      ],
      "title": "Last Terminated Exit Code",
      "type": "stat"
    },
    {
      "collapsed": false,
      "datasource": {
        "type": "datasource",
        "uid": "grafana"
      },
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 5
      },
      "id": 47,
      "panels": [],
      "targets": [
        {
          "datasource": {
            "type": "datasource",
            "uid": "grafana"
          },
          "refId": "A"
        }
      ],
      "title": "Resources",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "fixedColor": "blue",
            "mode": "fixed"
          },
          "decimals": 2,
          "mappings": [],
          "max": 1,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "#EAB839",
                "value": 60
              },
              {
                "color": "red",
                "value": 75
              }
            ]
          },
          "unit": "percentunit"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 3,
        "x": 0,
        "y": 6
      },
      "id": 39,
      "options": {
        "minVizHeight": 75,
        "minVizWidth": 75,
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "last"
          ],
          "fields": "",
          "values": false
        },
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "sizing": "auto"
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", cluster=\"$cluster\"}[$__rate_interval])) / sum(kube_pod_container_resource_requests{namespace=\"$namespace\", pod=~\"$pod\", resource=\"cpu\", job=~\"$job\", cluster=\"$cluster\"})",
          "instant": true,
          "interval": "$resolution",
          "legendFormat": "Requests",
          "refId": "A"
        }
      ],
      "title": "Total pod CPU Requests usage",
      "type": "gauge"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "decimals": 2,
          "mappings": [],
          "max": 1,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "#EAB839",
                "value": 60
              },
              {
                "color": "red",
                "value": 75
              }
            ]
          },
          "unit": "percentunit"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 3,
        "x": 3,
        "y": 6
      },
      "id": 48,
      "options": {
        "minVizHeight": 75,
        "minVizWidth": 75,
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "last"
          ],
          "fields": "",
          "values": false
        },
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "sizing": "auto"
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", cluster=\"$cluster\"}[$__rate_interval])) / sum(kube_pod_container_resource_limits{namespace=\"$namespace\", pod=~\"$pod\", resource=\"cpu\", job=~\"$job\", cluster=\"$cluster\"})",
          "instant": true,
          "interval": "$resolution",
          "legendFormat": "Limits",
          "refId": "A"
        }
      ],
      "title": "Total pod CPU Limits usage",
      "type": "gauge"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "fixedColor": "blue",
            "mode": "fixed"
          },
          "decimals": 2,
          "mappings": [],
          "max": 1,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "blue",
                "value": null
              },
              {
                "color": "#EAB839",
                "value": 80
              },
              {
                "color": "red",
                "value": 99
              }
            ]
          },
          "unit": "percentunit"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 3,
        "x": 6,
        "y": 6
      },
      "id": 40,
      "options": {
        "minVizHeight": 75,
        "minVizWidth": 75,
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "last"
          ],
          "fields": "",
          "values": false
        },
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "sizing": "auto"
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(container_memory_working_set_bytes{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", cluster=\"$cluster\"}) / sum(kube_pod_container_resource_requests{namespace=\"$namespace\", pod=~\"$pod\", resource=\"memory\", job=~\"$job\", cluster=\"$cluster\"})",
          "instant": true,
          "interval": "$resolution",
          "legendFormat": "Requests",
          "refId": "A"
        }
      ],
      "title": "Total pod RAM Requests usage",
      "type": "gauge"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "decimals": 2,
          "mappings": [],
          "max": 1,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "#EAB839",
                "value": 60
              },
              {
                "color": "red",
                "value": 75
              }
            ]
          },
          "unit": "percentunit"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 3,
        "x": 9,
        "y": 6
      },
      "id": 49,
      "options": {
        "minVizHeight": 75,
        "minVizWidth": 75,
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "last"
          ],
          "fields": "",
          "values": false
        },
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "sizing": "auto"
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(container_memory_working_set_bytes{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", cluster=\"$cluster\"}) / sum(kube_pod_container_resource_limits{namespace=\"$namespace\", pod=~\"$pod\", resource=\"memory\", job=~\"$job\", cluster=\"$cluster\"}) ",
          "instant": true,
          "interval": "$resolution",
          "legendFormat": "Limits",
          "refId": "B"
        }
      ],
      "title": "Total pod RAM Limits usage",
      "type": "gauge"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "custom": {
            "align": "auto",
            "cellOptions": {
              "type": "auto"
            },
            "filterable": false,
            "inspect": false,
            "minWidth": 100
          },
          "decimals": 4,
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "blue",
                "value": null
              }
            ]
          },
          "unit": "none"
        },
        "overrides": [
          {
            "matcher": {
              "id": "byName",
              "options": "Memory Requests"
            },
            "properties": [
              {
                "id": "unit",
                "value": "bytes"
              },
              {
                "id": "decimals",
                "value": 2
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "Memory Limits"
            },
            "properties": [
              {
                "id": "unit",
                "value": "bytes"
              },
              {
                "id": "decimals",
                "value": 2
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "Memory Used"
            },
            "properties": [
              {
                "id": "unit",
                "value": "bytes"
              },
              {
                "id": "decimals",
                "value": 2
              }
            ]
          }
        ]
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 6
      },
      "id": 38,
      "options": {
        "cellHeight": "sm",
        "footer": {
          "countRows": false,
          "fields": "",
          "reducer": [
            "sum"
          ],
          "show": false
        },
        "showHeader": true,
        "sortBy": []
      },
      "pluginVersion": "11.2.0",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(kube_pod_container_resource_requests{namespace=\"$namespace\", pod=~\"$pod\", resource=\"cpu\", job=~\"$job\", cluster=\"$cluster\"}) by (container)",
          "format": "table",
          "instant": true,
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "",
          "refId": "A"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(kube_pod_container_resource_limits{namespace=\"$namespace\", pod=~\"$pod\", resource=\"cpu\", job=~\"$job\", cluster=\"$cluster\"}) by (container)",
          "format": "table",
          "instant": true,
          "interval": "",
          "intervalFactor": 1,
          "legendFormat": "",
          "refId": "B"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(kube_pod_container_resource_requests{namespace=\"$namespace\", pod=~\"$pod\", resource=\"memory\", job=~\"$job\", cluster=\"$cluster\"}) by (container)",
          "format": "table",
          "instant": true,
          "interval": "",
          "legendFormat": "",
          "refId": "C"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(kube_pod_container_resource_limits{namespace=\"$namespace\", pod=~\"$pod\", resource=\"memory\", job=~\"$job\", cluster=\"$cluster\"}) by (container)",
          "format": "table",
          "instant": true,
          "interval": "",
          "legendFormat": "",
          "refId": "D"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", container!=\"\", cluster=\"$cluster\"}[$__rate_interval])) by (container)",
          "format": "table",
          "hide": false,
          "instant": true,
          "legendFormat": "__auto",
          "range": false,
          "refId": "E"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(container_memory_working_set_bytes{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", container!=\"\", cluster=\"$cluster\"}) by (container)",
          "format": "table",
          "hide": false,
          "instant": true,
          "range": false,
          "refId": "F"
        }
      ],
      "title": "Resources by container",
      "transformations": [
        {
          "id": "seriesToColumns",
          "options": {
            "byField": "container"
          }
        },
        {
          "id": "organize",
          "options": {
            "excludeByName": {
              "Time": true,
              "Time 1": true,
              "Time 2": true,
              "Time 4": true,
              "__name__": true,
              "__name__ 1": true,
              "__name__ 2": true,
              "__name__ 3": true,
              "__name__ 4": true,
              "container": false,
              "endpoint": true,
              "endpoint 2": true,
              "endpoint 3": true,
              "endpoint 4": true,
              "instance": true,
              "instance 2": true,
              "instance 3": true,
              "instance 4": true,
              "job": true,
              "job 2": true,
              "job 3": true,
              "job 4": true,
              "namespace": true,
              "namespace 2": true,
              "namespace 3": true,
              "namespace 4": true,
              "node": true,
              "node 2": true,
              "node 3": true,
              "node 4": true,
              "pod": true,
              "pod 2": true,
              "pod 3": true,
              "pod 4": true,
              "resource 1": true,
              "resource 2": true,
              "resource 3": true,
              "resource 4": true,
              "service": true,
              "service 2": true,
              "service 3": true,
              "service 4": true,
              "uid 1": true,
              "uid 2": true,
              "uid 3": true,
              "uid 4": true,
              "unit 1": true,
              "unit 2": true,
              "unit 3": true,
              "unit 4": true
            },
            "indexByName": {
              "Time 1": 7,
              "Time 2": 8,
              "Time 3": 9,
              "Time 4": 10,
              "Time 5": 11,
              "Time 6": 12,
              "Value #A": 2,
              "Value #B": 3,
              "Value #C": 5,
              "Value #D": 6,
              "Value #E": 1,
              "Value #F": 4,
              "container": 0
            },
            "renameByName": {
              "Value #A": "CPU Requests",
              "Value #B": "CPU Limits",
              "Value #C": "Memory Requests",
              "Value #D": "Memory Limits",
              "Value #E": "CPU Used",
              "Value #F": "Memory Used",
              "container": "Container"
            }
          }
        }
      ],
      "type": "table"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Percent",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "area"
            }
          },
          "mappings": [],
          "max": 1,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "red",
                "value": null
              },
              {
                "color": "yellow",
                "value": 20
              },
              {
                "color": "green",
                "value": 30
              },
              {
                "color": "yellow",
                "value": 70
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "percentunit"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 14
      },
      "id": 50,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "table",
          "placement": "right",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", cluster=\"$cluster\"}[$__rate_interval])) by (container) / sum(kube_pod_container_resource_requests{namespace=\"$namespace\", pod=~\"$pod\", resource=\"cpu\", job=~\"$job\", cluster=\"$cluster\"}) by (container)",
          "interval": "$resolution",
          "legendFormat": "{{ container }}  REQUESTS",
          "range": true,
          "refId": "A"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", cluster=\"$cluster\"}[$__rate_interval])) by (container) / sum(kube_pod_container_resource_limits{namespace=\"$namespace\", pod=~\"$pod\", resource=\"cpu\", job=~\"$job\", cluster=\"$cluster\"}) by (container)",
          "hide": false,
          "legendFormat": "{{ container }}  LIMITS",
          "range": true,
          "refId": "B"
        }
      ],
      "title": "CPU Usage / Requests & Limits by container",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "fixedColor": "blue",
            "mode": "thresholds"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Percent",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineStyle": {
              "fill": "solid"
            },
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "area"
            }
          },
          "mappings": [],
          "max": 1,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "red",
                "value": null
              },
              {
                "color": "yellow",
                "value": 20
              },
              {
                "color": "green",
                "value": 30
              },
              {
                "color": "#EAB839",
                "value": 70
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "percentunit"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 14
      },
      "id": 30,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "table",
          "placement": "right",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(container_memory_working_set_bytes{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", cluster=\"$cluster\"}) by (container) / sum(kube_pod_container_resource_requests{namespace=\"$namespace\", pod=~\"$pod\", resource=\"memory\", job=~\"$job\", cluster=\"$cluster\"}) by (container)",
          "interval": "",
          "legendFormat": "{{ container }} REQUESTS",
          "range": true,
          "refId": "A"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "expr": "sum(container_memory_working_set_bytes{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", cluster=\"$cluster\"}) by (container) / sum(kube_pod_container_resource_limits{namespace=\"$namespace\", pod=~\"$pod\", resource=\"memory\", job=~\"$job\", cluster=\"$cluster\"}) by (container)",
          "hide": false,
          "legendFormat": "{{ container }} LIMITS",
          "range": true,
          "refId": "B"
        }
      ],
      "title": "Memory Usage / Requests & Limits by container",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "CPU Cores",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "decimals": 4,
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "none"
        },
        "overrides": [
          {
            "matcher": {
              "id": "byName",
              "options": "limit"
            },
            "properties": [
              {
                "id": "color",
                "value": {
                  "fixedColor": "#F2495C",
                  "mode": "fixed"
                }
              },
              {
                "id": "custom.fillOpacity",
                "value": 0
              }
            ]
          }
        ]
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 22
      },
      "id": 29,
      "options": {
        "legend": {
          "calcs": [
            "min",
            "max",
            "mean"
          ],
          "displayMode": "table",
          "placement": "right",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", container!=\"\", cluster=\"$cluster\"}[$__rate_interval])) by (container)",
          "interval": "$resolution",
          "legendFormat": "{{ container }}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "CPU Usage by container",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Bytes",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineStyle": {
              "fill": "solid"
            },
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "decimals": 2,
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "bytes"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 22
      },
      "id": 51,
      "options": {
        "legend": {
          "calcs": [
            "min",
            "max",
            "mean"
          ],
          "displayMode": "table",
          "placement": "right",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(container_memory_working_set_bytes{namespace=\"$namespace\", pod=~\"$pod\", image!=\"\", container!=\"\", cluster=\"$cluster\"}) by (container)",
          "interval": "",
          "legendFormat": "{{ container }}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "Memory Usage by container",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "description": "",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "SECONDS",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineStyle": {
              "fill": "solid"
            },
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "decimals": 2,
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "s"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 30
      },
      "id": 59,
      "options": {
        "legend": {
          "calcs": [
            "min",
            "max",
            "mean"
          ],
          "displayMode": "table",
          "placement": "right",
          "showLegend": true,
          "sortBy": "Max",
          "sortDesc": true
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(rate(container_cpu_cfs_throttled_seconds_total{namespace=~\"$namespace\", pod=~\"$pod\", image!=\"\", container!=\"\", cluster=\"$cluster\"}[$__rate_interval])) by (container)",
          "interval": "$resolution",
          "legendFormat": "{{ container }}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "CPU Throttled seconds by container",
      "type": "timeseries"
    },
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 38
      },
      "id": 62,
      "panels": [],
      "title": "Kubernetes",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "fixedColor": "blue",
            "mode": "thresholds"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Percent",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineStyle": {
              "fill": "solid"
            },
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "area"
            }
          },
          "mappings": [],
          "max": 1,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "red",
                "value": null
              },
              {
                "color": "yellow",
                "value": 20
              },
              {
                "color": "green",
                "value": 30
              },
              {
                "color": "#EAB839",
                "value": 70
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 39
      },
      "id": 60,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "table",
          "placement": "right",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(increase(container_oom_events_total{namespace=\"${namespace}\", pod=\"${pod}\", container!=\"\", cluster=\"$cluster\"}[$__rate_interval])) by (container)",
          "interval": "",
          "legendFormat": "{{ container }}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "OOM Events by container",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "fixedColor": "blue",
            "mode": "thresholds"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Percent",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineStyle": {
              "fill": "solid"
            },
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "area"
            }
          },
          "mappings": [],
          "max": 1,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "red",
                "value": null
              },
              {
                "color": "yellow",
                "value": 20
              },
              {
                "color": "green",
                "value": 30
              },
              {
                "color": "#EAB839",
                "value": 70
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "none"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 39
      },
      "id": 61,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "table",
          "placement": "right",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(increase(kube_pod_container_status_restarts_total{namespace=~\"${namespace}\", pod=\"${pod}\", container!=\"\", job=~\"$job\", cluster=\"$cluster\"}[$__rate_interval])) by (container)",
          "interval": "",
          "legendFormat": "{{ container }}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "Container Restarts by container",
      "type": "timeseries"
    },
    {
      "collapsed": false,
      "datasource": {
        "type": "datasource",
        "uid": "grafana"
      },
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 47
      },
      "id": 45,
      "panels": [],
      "targets": [
        {
          "datasource": {
            "type": "datasource",
            "uid": "grafana"
          },
          "refId": "A"
        }
      ],
      "title": "Network",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "binBps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 48
      },
      "id": 31,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "none"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "exemplar": true,
          "expr": "sum(rate(container_network_receive_bytes_total{namespace=\"$namespace\", pod=~\"$pod\", cluster=\"$cluster\"}[$__rate_interval]))",
          "interval": "$resolution",
          "legendFormat": "Received",
          "refId": "A"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "exemplar": true,
          "expr": "- sum(rate(container_network_transmit_bytes_total{namespace=\"$namespace\", pod=~\"$pod\", cluster=\"$cluster\"}[$__rate_interval]))",
          "interval": "$resolution",
          "legendFormat": "Transmitted",
          "refId": "B"
        }
      ],
      "title": "Network - Bandwidth",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "pps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 48
      },
      "id": 34,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "none"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "exemplar": true,
          "expr": "sum(rate(container_network_receive_packets_total{namespace=\"$namespace\", pod=~\"$pod\", cluster=\"$cluster\"}[$__rate_interval]))",
          "interval": "$resolution",
          "legendFormat": "Received",
          "refId": "A"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "exemplar": true,
          "expr": "- sum(rate(container_network_transmit_packets_total{namespace=\"$namespace\", pod=~\"$pod\", cluster=\"$cluster\"}[$__rate_interval]))",
          "interval": "$resolution",
          "legendFormat": "Transmitted",
          "refId": "B"
        }
      ],
      "title": "Network - Packets Rate",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "pps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 56
      },
      "id": 36,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "none"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "exemplar": true,
          "expr": "sum(rate(container_network_receive_packets_dropped_total{namespace=\"$namespace\", pod=~\"$pod\", cluster=\"$cluster\"}[$__rate_interval]))",
          "interval": "$resolution",
          "legendFormat": "Received",
          "refId": "A"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "exemplar": true,
          "expr": "- sum(rate(container_network_transmit_packets_dropped_total{namespace=\"$namespace\", pod=~\"$pod\", cluster=\"$cluster\"}[$__rate_interval]))",
          "interval": "$resolution",
          "legendFormat": "Transmitted",
          "refId": "B"
        }
      ],
      "title": "Network - Packets Dropped",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${datasource}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "smooth",
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "pps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 56
      },
      "id": 37,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "none"
        }
      },
      "pluginVersion": "8.3.3",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "exemplar": true,
          "expr": "sum(rate(container_network_receive_errors_total{namespace=\"$namespace\", pod=~\"$pod\", cluster=\"$cluster\"}[$__rate_interval]))",
          "interval": "$resolution",
          "legendFormat": "Received",
          "refId": "A"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${datasource}"
          },
          "exemplar": true,
          "expr": "- sum(rate(container_network_transmit_errors_total{namespace=\"$namespace\", pod=~\"$pod\", cluster=\"$cluster\"}[$__rate_interval]))",
          "interval": "$resolution",
          "legendFormat": "Transmitted",
          "refId": "B"
        }
      ],
      "title": "Network - Errors",
      "type": "timeseries"
    }
  ],
  "refresh": "30s",
  "schemaVersion": 39,
  "tags": [
    "Kubernetes",
    "Prometheus",
    "KubePi"
  ],
  "templating": {
    "list": [
      {
        "current": {
          "selected": false,
          "text": "",
          "value": ""
        },
        "hide": 0,
        "includeAll": false,
        "multi": false,
        "name": "datasource",
        "options": [],
        "query": "prometheus",
        "queryValue": "",
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "type": "datasource"
      },
      {
        "current": {
          "isNone": true,
          "selected": false,
          "text": "None",
          "value": ""
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${datasource}"
        },
        "definition": "label_values(kube_node_info,cluster)",
        "hide": 0,
        "includeAll": false,
        "multi": false,
        "name": "cluster",
        "options": [],
        "query": {
          "qryType": 1,
          "query": "label_values(kube_node_info,cluster)",
          "refId": "PrometheusVariableQueryEditor-VariableQuery"
        },
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "type": "query"
      },
      {
        "current": {
          "selected": false,
          "text": "monitoring",
          "value": "monitoring"
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${datasource}"
        },
        "definition": "label_values(kube_pod_info{cluster=\"$cluster\"}, namespace)",
        "hide": 0,
        "includeAll": false,
        "multi": false,
        "name": "namespace",
        "options": [],
        "query": {
          "query": "label_values(kube_pod_info{cluster=\"$cluster\"}, namespace)",
          "refId": "Prometheus-namespace-Variable-Query"
        },
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "tagValuesQuery": "",
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
      {
        "current": {
          "selected": false,
          "text": "",
          "value": ""
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${datasource}"
        },
        "definition": "label_values(kube_pod_info{namespace=\"$namespace\", cluster=\"$cluster\"}, pod)",
        "hide": 0,
        "includeAll": true,
        "multi": true,
        "name": "pod",
        "options": [],
        "query": {
          "query": "label_values(kube_pod_info{namespace=\"$namespace\", cluster=\"$cluster\"}, pod)",
          "refId": "Prometheus-pod-Variable-Query"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "tagValuesQuery": "",
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
      {
        "current": {
          "selected": false,
          "text": "30s",
          "value": "30s"
        },
        "hide": 0,
        "includeAll": false,
        "multi": false,
        "name": "resolution",
        "options": [
          {
            "selected": false,
            "text": "1s",
            "value": "1s"
          },
          {
            "selected": false,
            "text": "15s",
            "value": "15s"
          },
          {
            "selected": true,
            "text": "30s",
            "value": "30s"
          },
          {
            "selected": false,
            "text": "1m",
            "value": "1m"
          },
          {
            "selected": false,
            "text": "3m",
            "value": "3m"
          },
          {
            "selected": false,
            "text": "5m",
            "value": "5m"
          }
        ],
        "query": "1s, 15s, 30s, 1m, 3m, 5m",
        "queryValue": "",
        "skipUrlSync": false,
        "type": "custom"
      },
      {
        "current": {
          "selected": false,
          "text": "kube-state-metrics",
          "value": "kube-state-metrics"
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${datasource}"
        },
        "definition": "label_values(kube_pod_info{namespace=\"$namespace\", cluster=\"$cluster\"},job)",
        "hide": 0,
        "includeAll": false,
        "multi": true,
        "name": "job",
        "options": [],
        "query": {
          "qryType": 1,
          "query": "label_values(kube_pod_info{namespace=\"$namespace\", cluster=\"$cluster\"},job)",
          "refId": "PrometheusVariableQueryEditor-VariableQuery"
        },
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "type": "query"
      }
    ]
  },
  "time": {
    "from": "now-1h",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "Kubernetes Pods Overview",
  "uid": "k8s_pod_overview_kubepi",
  "version": 32,
  "weekStart": "",
  "gnetId": 15760
}`

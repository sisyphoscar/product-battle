document.addEventListener("DOMContentLoaded", async () => {
    const brokerEndpoint = document.getElementById("scoreChart").dataset.brokerEndpoint;

    try {
        // 從 API 獲取資料
        const response = await fetch(`${brokerEndpoint}/api/widgets/product-score`);
        if (!response.ok) {
            throw new Error(`Failed to fetch data: ${response.status}`);
        }

        const widgetData = await response.json();

        // 提取資料
        const labels = widgetData.stats.map(stat => stat.productName);
        const data = widgetData.stats.map(stat => stat.score);

        // 使用 Chart.js 繪製直條圖
        const ctx = document.getElementById('scoreChart').getContext('2d');
        new Chart(ctx, {
            type: 'bar',
            data: {
                labels: labels,
                datasets: [{
                    label: 'Scores',
                    data: data,
                    backgroundColor: ['#4586F0', '#76D7C4', '#E74C3C', '#9B59B6'],
                    borderWidth: 1
                }]
            },
            options: {
                responsive: true,
                plugins: {
                    legend: {
                        display: false,
                        position: 'top'
                    },
                    tooltip: {
                        enabled: true
                    }
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        title: {
                            display: true,
                            text: 'Score'
                        }
                    },
                    x: {
                        title: {
                            display: true,
                            text: 'Product Name'
                        }
                    }
                }
            }
        });
    } catch (error) {
        console.error("Error loading chart data:", error);
        document.getElementById('scoreChart').outerHTML = `<p>Failed to load chart data. Please try again later.</p>`;
    }
});
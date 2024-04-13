const option = {
  series: [
    {
      type: 'gauge',
      startAngle: 90,
      endAngle: -270,
      pointer: {
        show: false
      },
      progress: {
        show: true,
        overlap: false,
        roundCap: true,
        clip: false
      },
      axisLine: {
        lineStyle: {
          width: 20
        }
      },
      splitLine: {
        show: false,
        distance: 0,
        length: 10
      },
      axisTick: {
        show: false
      },
      axisLabel: {
        show: false,
        distance: 50
      },
      data: [
        {
          value: 0,
          name: '服务器内存使用率',
          title: {
            offsetCenter: ['0%', '-20%']
          },
          detail: {
            valueAnimation: true,
            offsetCenter: ['0%', '20%']
          }
        }
      ],
      title: {
        fontSize: 28
      },
      detail: {
        width: 50,
        height: 20,
        fontSize: 30,
        color: 'inherit',
        formatter: '{value}%'
      }
    }
  ]
}
export default option

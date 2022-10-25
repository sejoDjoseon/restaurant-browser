import { useEffect } from 'react'

import { autorun } from 'mobx'

export const useAutorun = (action: Function) => {
  useEffect(() => {
    const disposer = autorun(() => {
      action()
    })

    return () => {
      disposer()
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])
}
